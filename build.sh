version=$1
sed "s/v.*/var Version=\"v$version\"/" b.txt > b.go
git add .
git commit -m .
git tag v$1
git push origin v$1
