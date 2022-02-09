find . -name '*.DS_Store' -type f -delete
cd ..
echo 'Start Add to GIT Repo'
git add Makefile
git commit -m 'auto commit'
git push
echo 'End Add to GIT Repo'
cd script/