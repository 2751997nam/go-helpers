folders=$(cd .. && ls -t -U | grep -v "system" | grep "service")
prefix="/helpers"
BASE_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
BASE_DIR=${BASE_DIR/%$prefix}
echo $BASE_DIR

for folder in ../*; do
    if [[ $folder == *"service"* ]]; then
        lsresult=$(ls $folder | grep "go.mod")
        if ls $folder | grep -q "go.mod"; then
            prefix=".."
            dir=${folder/#$prefix}
            echo $BASE_DIR$dir
            cd $BASE_DIR$dir && go get github.com/2751997nam/go-helpers@master && go mod tidy
        fi
    fi
done