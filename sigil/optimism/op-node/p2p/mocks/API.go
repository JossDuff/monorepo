// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	enode "github.com/ethereum/go-ethereum/p2p/enode"
	mock "github.com/stretchr/testify/mock"

	net "net"

	p2p "github.com/ethereum-optimism/optimism/op-node/p2p"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

type API_Expecter struct {
	mock *mock.Mock
}

func (_m *API) EXPECT() *API_Expecter {
	return &API_Expecter{mock: &_m.Mock}
}

// BlockAddr provides a mock function with given fields: ctx, ip
func (_m *API) BlockAddr(ctx context.Context, ip net.IP) error {
	ret := _m.Called(ctx, ip)

	if len(ret) == 0 {
		panic("no return value specified for BlockAddr")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, net.IP) error); ok {
		r0 = rf(ctx, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_BlockAddr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockAddr'
type API_BlockAddr_Call struct {
	*mock.Call
}

// BlockAddr is a helper method to define mock.On call
//   - ctx context.Context
//   - ip net.IP
func (_e *API_Expecter) BlockAddr(ctx interface{}, ip interface{}) *API_BlockAddr_Call {
	return &API_BlockAddr_Call{Call: _e.mock.On("BlockAddr", ctx, ip)}
}

func (_c *API_BlockAddr_Call) Run(run func(ctx context.Context, ip net.IP)) *API_BlockAddr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(net.IP))
	})
	return _c
}

func (_c *API_BlockAddr_Call) Return(_a0 error) *API_BlockAddr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_BlockAddr_Call) RunAndReturn(run func(context.Context, net.IP) error) *API_BlockAddr_Call {
	_c.Call.Return(run)
	return _c
}

// BlockPeer provides a mock function with given fields: ctx, p
func (_m *API) BlockPeer(ctx context.Context, p peer.ID) error {
	ret := _m.Called(ctx, p)

	if len(ret) == 0 {
		panic("no return value specified for BlockPeer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, peer.ID) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_BlockPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockPeer'
type API_BlockPeer_Call struct {
	*mock.Call
}

// BlockPeer is a helper method to define mock.On call
//   - ctx context.Context
//   - p peer.ID
func (_e *API_Expecter) BlockPeer(ctx interface{}, p interface{}) *API_BlockPeer_Call {
	return &API_BlockPeer_Call{Call: _e.mock.On("BlockPeer", ctx, p)}
}

func (_c *API_BlockPeer_Call) Run(run func(ctx context.Context, p peer.ID)) *API_BlockPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(peer.ID))
	})
	return _c
}

func (_c *API_BlockPeer_Call) Return(_a0 error) *API_BlockPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_BlockPeer_Call) RunAndReturn(run func(context.Context, peer.ID) error) *API_BlockPeer_Call {
	_c.Call.Return(run)
	return _c
}

// BlockSubnet provides a mock function with given fields: ctx, ipnet
func (_m *API) BlockSubnet(ctx context.Context, ipnet *net.IPNet) error {
	ret := _m.Called(ctx, ipnet)

	if len(ret) == 0 {
		panic("no return value specified for BlockSubnet")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *net.IPNet) error); ok {
		r0 = rf(ctx, ipnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_BlockSubnet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockSubnet'
type API_BlockSubnet_Call struct {
	*mock.Call
}

// BlockSubnet is a helper method to define mock.On call
//   - ctx context.Context
//   - ipnet *net.IPNet
func (_e *API_Expecter) BlockSubnet(ctx interface{}, ipnet interface{}) *API_BlockSubnet_Call {
	return &API_BlockSubnet_Call{Call: _e.mock.On("BlockSubnet", ctx, ipnet)}
}

func (_c *API_BlockSubnet_Call) Run(run func(ctx context.Context, ipnet *net.IPNet)) *API_BlockSubnet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*net.IPNet))
	})
	return _c
}

func (_c *API_BlockSubnet_Call) Return(_a0 error) *API_BlockSubnet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_BlockSubnet_Call) RunAndReturn(run func(context.Context, *net.IPNet) error) *API_BlockSubnet_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectPeer provides a mock function with given fields: ctx, addr
func (_m *API) ConnectPeer(ctx context.Context, addr string) error {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for ConnectPeer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_ConnectPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectPeer'
type API_ConnectPeer_Call struct {
	*mock.Call
}

// ConnectPeer is a helper method to define mock.On call
//   - ctx context.Context
//   - addr string
func (_e *API_Expecter) ConnectPeer(ctx interface{}, addr interface{}) *API_ConnectPeer_Call {
	return &API_ConnectPeer_Call{Call: _e.mock.On("ConnectPeer", ctx, addr)}
}

func (_c *API_ConnectPeer_Call) Run(run func(ctx context.Context, addr string)) *API_ConnectPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *API_ConnectPeer_Call) Return(_a0 error) *API_ConnectPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_ConnectPeer_Call) RunAndReturn(run func(context.Context, string) error) *API_ConnectPeer_Call {
	_c.Call.Return(run)
	return _c
}

// DisconnectPeer provides a mock function with given fields: ctx, id
func (_m *API) DisconnectPeer(ctx context.Context, id peer.ID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DisconnectPeer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, peer.ID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_DisconnectPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisconnectPeer'
type API_DisconnectPeer_Call struct {
	*mock.Call
}

// DisconnectPeer is a helper method to define mock.On call
//   - ctx context.Context
//   - id peer.ID
func (_e *API_Expecter) DisconnectPeer(ctx interface{}, id interface{}) *API_DisconnectPeer_Call {
	return &API_DisconnectPeer_Call{Call: _e.mock.On("DisconnectPeer", ctx, id)}
}

func (_c *API_DisconnectPeer_Call) Run(run func(ctx context.Context, id peer.ID)) *API_DisconnectPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(peer.ID))
	})
	return _c
}

func (_c *API_DisconnectPeer_Call) Return(_a0 error) *API_DisconnectPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_DisconnectPeer_Call) RunAndReturn(run func(context.Context, peer.ID) error) *API_DisconnectPeer_Call {
	_c.Call.Return(run)
	return _c
}

// DiscoveryTable provides a mock function with given fields: ctx
func (_m *API) DiscoveryTable(ctx context.Context) ([]*enode.Node, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DiscoveryTable")
	}

	var r0 []*enode.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*enode.Node, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*enode.Node); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*enode.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_DiscoveryTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DiscoveryTable'
type API_DiscoveryTable_Call struct {
	*mock.Call
}

// DiscoveryTable is a helper method to define mock.On call
//   - ctx context.Context
func (_e *API_Expecter) DiscoveryTable(ctx interface{}) *API_DiscoveryTable_Call {
	return &API_DiscoveryTable_Call{Call: _e.mock.On("DiscoveryTable", ctx)}
}

func (_c *API_DiscoveryTable_Call) Run(run func(ctx context.Context)) *API_DiscoveryTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *API_DiscoveryTable_Call) Return(_a0 []*enode.Node, _a1 error) *API_DiscoveryTable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_DiscoveryTable_Call) RunAndReturn(run func(context.Context) ([]*enode.Node, error)) *API_DiscoveryTable_Call {
	_c.Call.Return(run)
	return _c
}

// ListBlockedAddrs provides a mock function with given fields: ctx
func (_m *API) ListBlockedAddrs(ctx context.Context) ([]net.IP, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListBlockedAddrs")
	}

	var r0 []net.IP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]net.IP, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []net.IP); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]net.IP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_ListBlockedAddrs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBlockedAddrs'
type API_ListBlockedAddrs_Call struct {
	*mock.Call
}

// ListBlockedAddrs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *API_Expecter) ListBlockedAddrs(ctx interface{}) *API_ListBlockedAddrs_Call {
	return &API_ListBlockedAddrs_Call{Call: _e.mock.On("ListBlockedAddrs", ctx)}
}

func (_c *API_ListBlockedAddrs_Call) Run(run func(ctx context.Context)) *API_ListBlockedAddrs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *API_ListBlockedAddrs_Call) Return(_a0 []net.IP, _a1 error) *API_ListBlockedAddrs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_ListBlockedAddrs_Call) RunAndReturn(run func(context.Context) ([]net.IP, error)) *API_ListBlockedAddrs_Call {
	_c.Call.Return(run)
	return _c
}

// ListBlockedPeers provides a mock function with given fields: ctx
func (_m *API) ListBlockedPeers(ctx context.Context) ([]peer.ID, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListBlockedPeers")
	}

	var r0 []peer.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]peer.ID, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []peer.ID); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]peer.ID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_ListBlockedPeers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBlockedPeers'
type API_ListBlockedPeers_Call struct {
	*mock.Call
}

// ListBlockedPeers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *API_Expecter) ListBlockedPeers(ctx interface{}) *API_ListBlockedPeers_Call {
	return &API_ListBlockedPeers_Call{Call: _e.mock.On("ListBlockedPeers", ctx)}
}

func (_c *API_ListBlockedPeers_Call) Run(run func(ctx context.Context)) *API_ListBlockedPeers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *API_ListBlockedPeers_Call) Return(_a0 []peer.ID, _a1 error) *API_ListBlockedPeers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_ListBlockedPeers_Call) RunAndReturn(run func(context.Context) ([]peer.ID, error)) *API_ListBlockedPeers_Call {
	_c.Call.Return(run)
	return _c
}

// ListBlockedSubnets provides a mock function with given fields: ctx
func (_m *API) ListBlockedSubnets(ctx context.Context) ([]*net.IPNet, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListBlockedSubnets")
	}

	var r0 []*net.IPNet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*net.IPNet, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*net.IPNet); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*net.IPNet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_ListBlockedSubnets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBlockedSubnets'
type API_ListBlockedSubnets_Call struct {
	*mock.Call
}

// ListBlockedSubnets is a helper method to define mock.On call
//   - ctx context.Context
func (_e *API_Expecter) ListBlockedSubnets(ctx interface{}) *API_ListBlockedSubnets_Call {
	return &API_ListBlockedSubnets_Call{Call: _e.mock.On("ListBlockedSubnets", ctx)}
}

func (_c *API_ListBlockedSubnets_Call) Run(run func(ctx context.Context)) *API_ListBlockedSubnets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *API_ListBlockedSubnets_Call) Return(_a0 []*net.IPNet, _a1 error) *API_ListBlockedSubnets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_ListBlockedSubnets_Call) RunAndReturn(run func(context.Context) ([]*net.IPNet, error)) *API_ListBlockedSubnets_Call {
	_c.Call.Return(run)
	return _c
}

// PeerStats provides a mock function with given fields: ctx
func (_m *API) PeerStats(ctx context.Context) (*p2p.PeerStats, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for PeerStats")
	}

	var r0 *p2p.PeerStats
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*p2p.PeerStats, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *p2p.PeerStats); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.PeerStats)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_PeerStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PeerStats'
type API_PeerStats_Call struct {
	*mock.Call
}

// PeerStats is a helper method to define mock.On call
//   - ctx context.Context
func (_e *API_Expecter) PeerStats(ctx interface{}) *API_PeerStats_Call {
	return &API_PeerStats_Call{Call: _e.mock.On("PeerStats", ctx)}
}

func (_c *API_PeerStats_Call) Run(run func(ctx context.Context)) *API_PeerStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *API_PeerStats_Call) Return(_a0 *p2p.PeerStats, _a1 error) *API_PeerStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_PeerStats_Call) RunAndReturn(run func(context.Context) (*p2p.PeerStats, error)) *API_PeerStats_Call {
	_c.Call.Return(run)
	return _c
}

// Peers provides a mock function with given fields: ctx, connected
func (_m *API) Peers(ctx context.Context, connected bool) (*p2p.PeerDump, error) {
	ret := _m.Called(ctx, connected)

	if len(ret) == 0 {
		panic("no return value specified for Peers")
	}

	var r0 *p2p.PeerDump
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) (*p2p.PeerDump, error)); ok {
		return rf(ctx, connected)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool) *p2p.PeerDump); ok {
		r0 = rf(ctx, connected)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.PeerDump)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, connected)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_Peers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Peers'
type API_Peers_Call struct {
	*mock.Call
}

// Peers is a helper method to define mock.On call
//   - ctx context.Context
//   - connected bool
func (_e *API_Expecter) Peers(ctx interface{}, connected interface{}) *API_Peers_Call {
	return &API_Peers_Call{Call: _e.mock.On("Peers", ctx, connected)}
}

func (_c *API_Peers_Call) Run(run func(ctx context.Context, connected bool)) *API_Peers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool))
	})
	return _c
}

func (_c *API_Peers_Call) Return(_a0 *p2p.PeerDump, _a1 error) *API_Peers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_Peers_Call) RunAndReturn(run func(context.Context, bool) (*p2p.PeerDump, error)) *API_Peers_Call {
	_c.Call.Return(run)
	return _c
}

// ProtectPeer provides a mock function with given fields: ctx, p
func (_m *API) ProtectPeer(ctx context.Context, p peer.ID) error {
	ret := _m.Called(ctx, p)

	if len(ret) == 0 {
		panic("no return value specified for ProtectPeer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, peer.ID) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_ProtectPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProtectPeer'
type API_ProtectPeer_Call struct {
	*mock.Call
}

// ProtectPeer is a helper method to define mock.On call
//   - ctx context.Context
//   - p peer.ID
func (_e *API_Expecter) ProtectPeer(ctx interface{}, p interface{}) *API_ProtectPeer_Call {
	return &API_ProtectPeer_Call{Call: _e.mock.On("ProtectPeer", ctx, p)}
}

func (_c *API_ProtectPeer_Call) Run(run func(ctx context.Context, p peer.ID)) *API_ProtectPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(peer.ID))
	})
	return _c
}

func (_c *API_ProtectPeer_Call) Return(_a0 error) *API_ProtectPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_ProtectPeer_Call) RunAndReturn(run func(context.Context, peer.ID) error) *API_ProtectPeer_Call {
	_c.Call.Return(run)
	return _c
}

// Self provides a mock function with given fields: ctx
func (_m *API) Self(ctx context.Context) (*p2p.PeerInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Self")
	}

	var r0 *p2p.PeerInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*p2p.PeerInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *p2p.PeerInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.PeerInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_Self_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Self'
type API_Self_Call struct {
	*mock.Call
}

// Self is a helper method to define mock.On call
//   - ctx context.Context
func (_e *API_Expecter) Self(ctx interface{}) *API_Self_Call {
	return &API_Self_Call{Call: _e.mock.On("Self", ctx)}
}

func (_c *API_Self_Call) Run(run func(ctx context.Context)) *API_Self_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *API_Self_Call) Return(_a0 *p2p.PeerInfo, _a1 error) *API_Self_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_Self_Call) RunAndReturn(run func(context.Context) (*p2p.PeerInfo, error)) *API_Self_Call {
	_c.Call.Return(run)
	return _c
}

// UnblockAddr provides a mock function with given fields: ctx, ip
func (_m *API) UnblockAddr(ctx context.Context, ip net.IP) error {
	ret := _m.Called(ctx, ip)

	if len(ret) == 0 {
		panic("no return value specified for UnblockAddr")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, net.IP) error); ok {
		r0 = rf(ctx, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_UnblockAddr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnblockAddr'
type API_UnblockAddr_Call struct {
	*mock.Call
}

// UnblockAddr is a helper method to define mock.On call
//   - ctx context.Context
//   - ip net.IP
func (_e *API_Expecter) UnblockAddr(ctx interface{}, ip interface{}) *API_UnblockAddr_Call {
	return &API_UnblockAddr_Call{Call: _e.mock.On("UnblockAddr", ctx, ip)}
}

func (_c *API_UnblockAddr_Call) Run(run func(ctx context.Context, ip net.IP)) *API_UnblockAddr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(net.IP))
	})
	return _c
}

func (_c *API_UnblockAddr_Call) Return(_a0 error) *API_UnblockAddr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_UnblockAddr_Call) RunAndReturn(run func(context.Context, net.IP) error) *API_UnblockAddr_Call {
	_c.Call.Return(run)
	return _c
}

// UnblockPeer provides a mock function with given fields: ctx, p
func (_m *API) UnblockPeer(ctx context.Context, p peer.ID) error {
	ret := _m.Called(ctx, p)

	if len(ret) == 0 {
		panic("no return value specified for UnblockPeer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, peer.ID) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_UnblockPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnblockPeer'
type API_UnblockPeer_Call struct {
	*mock.Call
}

// UnblockPeer is a helper method to define mock.On call
//   - ctx context.Context
//   - p peer.ID
func (_e *API_Expecter) UnblockPeer(ctx interface{}, p interface{}) *API_UnblockPeer_Call {
	return &API_UnblockPeer_Call{Call: _e.mock.On("UnblockPeer", ctx, p)}
}

func (_c *API_UnblockPeer_Call) Run(run func(ctx context.Context, p peer.ID)) *API_UnblockPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(peer.ID))
	})
	return _c
}

func (_c *API_UnblockPeer_Call) Return(_a0 error) *API_UnblockPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_UnblockPeer_Call) RunAndReturn(run func(context.Context, peer.ID) error) *API_UnblockPeer_Call {
	_c.Call.Return(run)
	return _c
}

// UnblockSubnet provides a mock function with given fields: ctx, ipnet
func (_m *API) UnblockSubnet(ctx context.Context, ipnet *net.IPNet) error {
	ret := _m.Called(ctx, ipnet)

	if len(ret) == 0 {
		panic("no return value specified for UnblockSubnet")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *net.IPNet) error); ok {
		r0 = rf(ctx, ipnet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_UnblockSubnet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnblockSubnet'
type API_UnblockSubnet_Call struct {
	*mock.Call
}

// UnblockSubnet is a helper method to define mock.On call
//   - ctx context.Context
//   - ipnet *net.IPNet
func (_e *API_Expecter) UnblockSubnet(ctx interface{}, ipnet interface{}) *API_UnblockSubnet_Call {
	return &API_UnblockSubnet_Call{Call: _e.mock.On("UnblockSubnet", ctx, ipnet)}
}

func (_c *API_UnblockSubnet_Call) Run(run func(ctx context.Context, ipnet *net.IPNet)) *API_UnblockSubnet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*net.IPNet))
	})
	return _c
}

func (_c *API_UnblockSubnet_Call) Return(_a0 error) *API_UnblockSubnet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_UnblockSubnet_Call) RunAndReturn(run func(context.Context, *net.IPNet) error) *API_UnblockSubnet_Call {
	_c.Call.Return(run)
	return _c
}

// UnprotectPeer provides a mock function with given fields: ctx, p
func (_m *API) UnprotectPeer(ctx context.Context, p peer.ID) error {
	ret := _m.Called(ctx, p)

	if len(ret) == 0 {
		panic("no return value specified for UnprotectPeer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, peer.ID) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_UnprotectPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnprotectPeer'
type API_UnprotectPeer_Call struct {
	*mock.Call
}

// UnprotectPeer is a helper method to define mock.On call
//   - ctx context.Context
//   - p peer.ID
func (_e *API_Expecter) UnprotectPeer(ctx interface{}, p interface{}) *API_UnprotectPeer_Call {
	return &API_UnprotectPeer_Call{Call: _e.mock.On("UnprotectPeer", ctx, p)}
}

func (_c *API_UnprotectPeer_Call) Run(run func(ctx context.Context, p peer.ID)) *API_UnprotectPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(peer.ID))
	})
	return _c
}

func (_c *API_UnprotectPeer_Call) Return(_a0 error) *API_UnprotectPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_UnprotectPeer_Call) RunAndReturn(run func(context.Context, peer.ID) error) *API_UnprotectPeer_Call {
	_c.Call.Return(run)
	return _c
}

// NewAPI creates a new instance of API. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *API {
	mock := &API{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}