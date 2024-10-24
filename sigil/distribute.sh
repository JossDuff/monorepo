./build
# nodea
echo "copying binary to nodea"
scp target/debug/sigil root@142.93.53.125:~/
# nodeb
echo "copying binary to nodeb"
scp target/debug/sigil root@142.93.2.49:~/
echo "\ndone! :)"
