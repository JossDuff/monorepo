package config

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"path"
	"slices"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer"
	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/artifacts"
	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/inspect"
	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/pipeline"
	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/state"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/exp/maps"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-chain-ops/foundry"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	op_service "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"

	_ "embed"
)

// legacy geth log levels - the geth command line --verbosity flag wasn't
// migrated to use slog's numerical levels.
const (
	LegacyLevelCrit = iota
	LegacyLevelError
	LegacyLevelWarn
	LegacyLevelInfo
	LegacyLevelDebug
	LegacyLevelTrace
)

type AllocType string

const (
	AllocTypeStandard AllocType = "standard"
	AllocTypeAltDA    AllocType = "alt-da"
	AllocTypeL2OO     AllocType = "l2oo"
	AllocTypeMTCannon AllocType = "mt-cannon"

	DefaultAllocType = AllocTypeStandard
)

func (a AllocType) Check() error {
	if !slices.Contains(allocTypes, a) {
		return fmt.Errorf("unknown alloc type: %q", a)
	}
	return nil
}

func (a AllocType) UsesProofs() bool {
	switch a {
	case AllocTypeStandard, AllocTypeMTCannon, AllocTypeAltDA:
		return true
	default:
		return false
	}
}

var allocTypes = []AllocType{AllocTypeStandard, AllocTypeAltDA, AllocTypeL2OO, AllocTypeMTCannon}

var (
	// All of the following variables are set in the init function
	// and read from JSON files on disk that are generated by the
	// foundry deploy script. These are globally exported to be used
	// in end to end tests.

	// L1Allocs represents the L1 genesis block state.
	l1AllocsByType = make(map[AllocType]*foundry.ForgeAllocs)
	// L1Deployments maps contract names to accounts in the L1
	// genesis block state.
	l1DeploymentsByType = make(map[AllocType]*genesis.L1Deployments)
	// l2Allocs represents the L2 allocs, by hardfork/mode (e.g. delta, ecotone, interop, other)
	l2AllocsByType = make(map[AllocType]genesis.L2AllocsModeMap)
	// DeployConfig represents the deploy config used by the system.
	deployConfigsByType = make(map[AllocType]*genesis.DeployConfig)
	// EthNodeVerbosity is the (legacy geth) level of verbosity to output
	EthNodeVerbosity int = 3

	// mtx is a lock to protect the above variables
	mtx sync.RWMutex
)

func L1Allocs(allocType AllocType) *foundry.ForgeAllocs {
	mtx.RLock()
	defer mtx.RUnlock()
	allocs, ok := l1AllocsByType[allocType]
	if !ok {
		panic(fmt.Errorf("unknown L1 alloc type: %q", allocType))
	}
	return allocs.Copy()
}

func L1Deployments(allocType AllocType) *genesis.L1Deployments {
	mtx.RLock()
	defer mtx.RUnlock()
	deployments, ok := l1DeploymentsByType[allocType]
	if !ok {
		panic(fmt.Errorf("unknown L1 deployments type: %q", allocType))
	}
	return deployments.Copy()
}

func L2Allocs(allocType AllocType, mode genesis.L2AllocsMode) *foundry.ForgeAllocs {
	mtx.RLock()
	defer mtx.RUnlock()
	allocsByType, ok := l2AllocsByType[allocType]
	if !ok {
		panic(fmt.Errorf("unknown L2 alloc type: %q", allocType))
	}

	allocs, ok := allocsByType[mode]
	if !ok {
		panic(fmt.Errorf("unknown L2 allocs mode: %q", mode))
	}
	return allocs.Copy()
}

func DeployConfig(allocType AllocType) *genesis.DeployConfig {
	mtx.RLock()
	defer mtx.RUnlock()
	dc, ok := deployConfigsByType[allocType]
	if !ok {
		panic(fmt.Errorf("unknown deploy config type: %q", allocType))
	}
	return dc.Copy()
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	root, err := op_service.FindMonorepoRoot(cwd)
	if err != nil {
		panic(err)
	}

	// Setup global logger
	lvl := log.FromLegacyLevel(EthNodeVerbosity)
	var handler slog.Handler
	var errHandler slog.Handler
	if lvl > log.LevelCrit {
		handler = log.DiscardHandler()
		errHandler = log.DiscardHandler()
	} else {
		if lvl < log.LevelTrace { // clip to trace level
			lvl = log.LevelTrace
		}
		// We cannot attach a testlog logger,
		// because the global logger is shared between different independent parallel tests.
		// Tests that write to a testlogger of another finished test fail.
		handler = oplog.NewLogHandler(os.Stdout, oplog.CLIConfig{
			Level:  lvl,
			Color:  false, // some CI logs do not handle colors well
			Format: oplog.FormatTerminal,
		})

		errHandler = oplog.NewLogHandler(os.Stderr, oplog.CLIConfig{
			Level:  log.LevelError,
			Color:  false,
			Format: oplog.FormatTerminal,
		})
	}

	// Start at warning level since alloc generation is heavy on the logs,
	// which reduces CI performance.
	oplog.SetGlobalLogHandler(errHandler)

	for _, allocType := range allocTypes {
		if allocType == AllocTypeL2OO {
			continue
		}

		initAllocType(root, allocType)
	}

	configPath := path.Join(root, "op-e2e", "config")
	forks := []genesis.L2AllocsMode{
		genesis.L2AllocsHolocene,
		genesis.L2AllocsGranite,
		genesis.L2AllocsFjord,
		genesis.L2AllocsEcotone,
		genesis.L2AllocsDelta,
	}

	var l2OOAllocsL1 foundry.ForgeAllocs
	decompressGzipJSON(path.Join(configPath, "allocs-l1.json.gz"), &l2OOAllocsL1)
	l1AllocsByType[AllocTypeL2OO] = &l2OOAllocsL1

	var l2OOAddresses genesis.L1Deployments
	decompressGzipJSON(path.Join(configPath, "addresses.json.gz"), &l2OOAddresses)
	l1DeploymentsByType[AllocTypeL2OO] = &l2OOAddresses

	l2OODC := DeployConfig(AllocTypeStandard)
	l2OODC.SetDeployments(&l2OOAddresses)
	deployConfigsByType[AllocTypeL2OO] = l2OODC

	l2AllocsByType[AllocTypeL2OO] = genesis.L2AllocsModeMap{}
	var wg sync.WaitGroup
	for _, fork := range forks {
		wg.Add(1)
		go func(fork genesis.L2AllocsMode) {
			defer wg.Done()
			var l2OOAllocsL2 foundry.ForgeAllocs
			decompressGzipJSON(path.Join(configPath, fmt.Sprintf("allocs-l2-%s.json.gz", fork)), &l2OOAllocsL2)
			mtx.Lock()
			l2AllocsByType[AllocTypeL2OO][fork] = &l2OOAllocsL2
			mtx.Unlock()
		}(fork)
	}
	wg.Wait()

	// Use regular level going forward.
	oplog.SetGlobalLogHandler(handler)
}

func initAllocType(root string, allocType AllocType) {
	artifactsPath := path.Join(root, "packages", "contracts-bedrock", "forge-artifacts")
	if err := ensureDir(artifactsPath); err != nil {
		panic(fmt.Errorf("invalid artifacts path: %w", err))
	}

	loc, err := artifacts.NewFileLocator(artifactsPath)
	if err != nil {
		panic(fmt.Errorf("failed to create artifacts locator: %w", err))
	}

	lgr := log.New()

	allocModes := []genesis.L2AllocsMode{
		genesis.L2AllocsHolocene,
		genesis.L2AllocsGranite,
		genesis.L2AllocsFjord,
		genesis.L2AllocsEcotone,
		genesis.L2AllocsDelta,
	}

	l2Alloc := make(map[genesis.L2AllocsMode]*foundry.ForgeAllocs)
	var wg sync.WaitGroup

	// Corresponds with the Deployer address in cfg.secrets
	pk, err := crypto.HexToECDSA("7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6")
	if err != nil {
		panic(fmt.Errorf("failed to parse private key: %w", err))
	}
	deployerAddr := crypto.PubkeyToAddress(pk.PublicKey)
	lgr.Info("deployer address", "address", deployerAddr.Hex())

	for _, mode := range allocModes {
		wg.Add(1)
		go func(mode genesis.L2AllocsMode) {
			defer wg.Done()

			intent := defaultIntent(root, loc, deployerAddr, allocType)
			if allocType == AllocTypeAltDA {
				intent.Chains[0].DangerousAltDAConfig = genesis.AltDADeployConfig{
					UseAltDA:                   true,
					DACommitmentType:           "KeccakCommitment",
					DAChallengeWindow:          16,
					DAResolveWindow:            16,
					DABondSize:                 1000000,
					DAResolverRefundPercentage: 0,
				}
			}

			baseUpgradeSchedule := map[string]any{
				"l2GenesisRegolithTimeOffset": nil,
				"l2GenesisCanyonTimeOffset":   nil,
				"l2GenesisDeltaTimeOffset":    nil,
				"l2GenesisEcotoneTimeOffset":  nil,
				"l2GenesisFjordTimeOffset":    nil,
				"l2GenesisGraniteTimeOffset":  nil,
				"l2GenesisHoloceneTimeOffset": nil,
			}

			upgradeSchedule := new(genesis.UpgradeScheduleDeployConfig)
			upgradeSchedule.ActivateForkAtGenesis(rollup.ForkName(mode))
			upgradeOverridesJSON, err := json.Marshal(upgradeSchedule)
			if err != nil {
				panic(fmt.Errorf("failed to marshal upgrade schedule: %w", err))
			}
			var upgradeOverrides map[string]any
			if err := json.Unmarshal(upgradeOverridesJSON, &upgradeOverrides); err != nil {
				panic(fmt.Errorf("failed to unmarshal upgrade schedule: %w", err))
			}
			maps.Copy(baseUpgradeSchedule, upgradeOverrides)
			maps.Copy(intent.GlobalDeployOverrides, baseUpgradeSchedule)

			st := &state.State{
				Version: 1,
			}

			if err := deployer.ApplyPipeline(
				context.Background(),
				deployer.ApplyPipelineOpts{
					DeploymentTarget:   deployer.DeploymentTargetGenesis,
					L1RPCUrl:           "",
					DeployerPrivateKey: pk,
					Intent:             intent,
					State:              st,
					Logger:             lgr,
					StateWriter:        pipeline.NoopStateWriter(),
				},
			); err != nil {
				panic(fmt.Errorf("failed to apply pipeline: %w", err))
			}

			mtx.Lock()
			l2Alloc[mode] = st.Chains[0].Allocs.Data
			mtx.Unlock()

			// This needs to be updated whenever the latest hardfork is changed.
			if mode == genesis.L2AllocsGranite {
				dc, err := inspect.DeployConfig(st, intent.Chains[0].ID)
				if err != nil {
					panic(fmt.Errorf("failed to inspect deploy config: %w", err))
				}

				l1Contracts, err := inspect.L1(st, intent.Chains[0].ID)
				if err != nil {
					panic(fmt.Errorf("failed to inspect L1: %w", err))
				}
				l1Deployments := l1Contracts.AsL1Deployments()

				// Set the L1 genesis block timestamp to now
				dc.L1GenesisBlockTimestamp = hexutil.Uint64(time.Now().Unix())
				dc.FundDevAccounts = true
				// Speed up the in memory tests
				dc.L1BlockTime = 2
				dc.L2BlockTime = 1
				dc.SetDeployments(l1Deployments)
				mtx.Lock()
				deployConfigsByType[allocType] = dc
				l1AllocsByType[allocType] = st.L1StateDump.Data
				l1DeploymentsByType[allocType] = l1Deployments
				mtx.Unlock()
			}
		}(mode)
	}
	wg.Wait()
	l2AllocsByType[allocType] = l2Alloc
}

func defaultIntent(root string, loc *artifacts.Locator, deployer common.Address, allocType AllocType) *state.Intent {
	defaultPrestate := common.HexToHash("0x03c7ae758795765c6664a5d39bf63841c71ff191e9189522bad8ebff5d4eca98")
	genesisOutputRoot := common.HexToHash("0xDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEF")
	return &state.Intent{
		ConfigType: state.IntentConfigTypeCustom,
		L1ChainID:  900,
		SuperchainRoles: &state.SuperchainRoles{
			ProxyAdminOwner:       deployer,
			ProtocolVersionsOwner: deployer,
			Guardian:              deployer,
		},
		FundDevAccounts:    true,
		L1ContractsLocator: loc,
		L2ContractsLocator: loc,
		GlobalDeployOverrides: map[string]any{
			"maxSequencerDrift":                        300,
			"sequencerWindowSize":                      200,
			"channelTimeout":                           120,
			"l2OutputOracleSubmissionInterval":         10,
			"l2OutputOracleStartingTimestamp":          0,
			"l2OutputOracleProposer":                   "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
			"l2OutputOracleChallenger":                 "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
			"l2GenesisBlockGasLimit":                   "0x1c9c380",
			"l1BlockTime":                              6,
			"baseFeeVaultMinimumWithdrawalAmount":      "0x8ac7230489e80000",
			"l1FeeVaultMinimumWithdrawalAmount":        "0x8ac7230489e80000",
			"sequencerFeeVaultMinimumWithdrawalAmount": "0x8ac7230489e80000",
			"baseFeeVaultWithdrawalNetwork":            0,
			"l1FeeVaultWithdrawalNetwork":              0,
			"sequencerFeeVaultWithdrawalNetwork":       0,
			"finalizationPeriodSeconds":                2,
			"l2GenesisBlockBaseFeePerGas":              "0x1",
			"gasPriceOracleOverhead":                   2100,
			"gasPriceOracleScalar":                     1000000,
			"gasPriceOracleBaseFeeScalar":              1368,
			"gasPriceOracleBlobBaseFeeScalar":          810949,
			"l1CancunTimeOffset":                       "0x0",
			"faultGameAbsolutePrestate":                defaultPrestate.Hex(),
			"faultGameMaxDepth":                        50,
			"faultGameClockExtension":                  0,
			"faultGameMaxClockDuration":                1200,
			"faultGameGenesisBlock":                    0,
			"faultGameGenesisOutputRoot":               genesisOutputRoot.Hex(),
			"faultGameSplitDepth":                      14,
			"dangerouslyAllowCustomDisputeParameters":  true,
			"faultGameWithdrawalDelay":                 604800,
			"preimageOracleMinProposalSize":            10000,
			"preimageOracleChallengePeriod":            120,
			"proofMaturityDelaySeconds":                12,
			"disputeGameFinalityDelaySeconds":          6,
		},
		Chains: []*state.ChainIntent{
			{
				ID:                         common.BigToHash(big.NewInt(901)),
				BaseFeeVaultRecipient:      common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"),
				L1FeeVaultRecipient:        common.HexToAddress("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f"),
				SequencerFeeVaultRecipient: common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"),
				Eip1559Denominator:         250,
				Eip1559DenominatorCanyon:   250,
				Eip1559Elasticity:          6,
				Roles: state.ChainRoles{
					// Use deployer as L1PAO to deploy additional dispute impls
					L1ProxyAdminOwner: deployer,
					L2ProxyAdminOwner: deployer,
					SystemConfigOwner: deployer,
					UnsafeBlockSigner: common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"),
					Batcher:           common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
					Proposer:          common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
					Challenger:        common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
				},
				AdditionalDisputeGames: []state.AdditionalDisputeGame{
					{
						ChainProofParams: state.ChainProofParams{
							// Fast game
							DisputeGameType:         254,
							DisputeAbsolutePrestate: defaultPrestate,
							DisputeMaxGameDepth:     14 + 3 + 1,
							DisputeSplitDepth:       14,
							DisputeClockExtension:   0,
							DisputeMaxClockDuration: 0,
						},
						VMType:                       state.VMTypeAlphabet,
						UseCustomOracle:              true,
						OracleMinProposalSize:        10000,
						OracleChallengePeriodSeconds: 0,
						MakeRespected:                true,
					},
					{
						ChainProofParams: state.ChainProofParams{
							// Alphabet game
							DisputeGameType:         255,
							DisputeAbsolutePrestate: defaultPrestate,
							DisputeMaxGameDepth:     14 + 3 + 1,
							DisputeSplitDepth:       14,
							DisputeClockExtension:   0,
							DisputeMaxClockDuration: 1200,
						},
						VMType: state.VMTypeAlphabet,
					},
					{
						ChainProofParams: state.ChainProofParams{
							DisputeGameType:         0,
							DisputeAbsolutePrestate: cannonPrestate(root, allocType),
							DisputeMaxGameDepth:     50,
							DisputeSplitDepth:       14,
							DisputeClockExtension:   0,
							DisputeMaxClockDuration: 1200,
						},
						VMType: cannonVMType(allocType),
					},
				},
			},
		},
	}
}

func ensureDir(dirPath string) error {
	stat, err := os.Stat(dirPath)
	if err != nil {
		return fmt.Errorf("failed to stat path: %w", err)
	}
	if !stat.IsDir() {
		return fmt.Errorf("path is not a directory")
	}
	return nil
}

func decompressGzipJSON(p string, thing any) {
	f, err := os.Open(p)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %w", err))
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		panic(fmt.Errorf("failed to create gzip reader: %w", err))
	}
	defer gzr.Close()
	if err := json.NewDecoder(gzr).Decode(thing); err != nil {
		panic(fmt.Errorf("failed to read gzip data: %w", err))
	}
}

func cannonVMType(allocType AllocType) state.VMType {
	if allocType == AllocTypeMTCannon {
		return state.VMTypeCannon2
	}
	return state.VMTypeCannon1
}

type prestateFile struct {
	Pre string `json:"pre"`
}

var cannonPrestateMT common.Hash
var cannonPrestateST common.Hash
var cannonPrestateMTOnce sync.Once
var cannonPrestateSTOnce sync.Once

func cannonPrestate(monorepoRoot string, allocType AllocType) common.Hash {
	var filename string

	var once *sync.Once
	var cacheVar *common.Hash
	if cannonVMType(allocType) == state.VMTypeCannon1 {
		filename = "prestate-proof.json"
		once = &cannonPrestateSTOnce
		cacheVar = &cannonPrestateST
	} else {
		filename = "prestate-proof-mt.json"
		once = &cannonPrestateMTOnce
		cacheVar = &cannonPrestateMT
	}

	once.Do(func() {
		f, err := os.Open(path.Join(monorepoRoot, "op-program", "bin", filename))
		if err != nil {
			log.Warn("error opening prestate file", "err", err)
			return
		}
		defer f.Close()

		var prestate prestateFile
		dec := json.NewDecoder(f)
		if err := dec.Decode(&prestate); err != nil {
			log.Warn("error decoding prestate file", "err", err)
			return
		}

		*cacheVar = common.HexToHash(prestate.Pre)
	})

	return *cacheVar
}