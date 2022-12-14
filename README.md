# gitOperationsDemo

https://github.com/nikkula/gitOperationsDemo

https://github.com/nikkula/gitOperationsDemo.git

uttarkumar.nikkula@peerislands.io


Git Commands:
-------------
1.  git config  -> Configure the UserName & Email Address
2.  git init  -> Initialize a local git repository
3.  git add -> Add one or more files to staging area
4.  git diff  -> View the changes made to the file
5.  git commit  -> Commit the changes to head but not to the remote repository
6.  git reset -> Undo local changes to the state of a Git repo
7.  git status  -> Displays the state of the working directory and staging area
8.  git merge -> Merge a branch into an active branch
9.  git push  -> Upload content from local repository to a remote repository
10.  git pull  -> Fetch and downaload content from a remote repository

Step 1:
git config --global user.name "xxxx"
git config -global user.email xxxx@xxx.com
git config --list
//pwd -> Present working directory
//After adding new files to the folder then run status command
git status
git add .
git status
git commit -m "Commit Message any text"
git log

//Push the local changes to git Repo
git remote add origin https://demourl.com/repositoryname.git
git remote -v
git push -u origin master
--------------------------
Step 2:
//Downloading the git remote repository code to my local repository

git pull https://github.com/nikkula/gitOperationsDemo.git

