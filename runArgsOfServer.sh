
# i=1
# go run main.go ":600$i"
for i in {1..5}
do
   a='/home/spurge/goNetworks/simpleServerArgs'
   cd $a;
   go run main.go ":500$i"
done

