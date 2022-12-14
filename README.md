# gitOperationsDemo
All these commands has to be run in GITBASH command window.

Reference URLS:
  1.  https://www.youtube.com/watch?v=QJ0iUNe27c8
  2.  https://www.youtube.com/watch?v=z8CYDyFqzp0
  
Error Hint : https://www.youtube.com/watch?v=WCTzGrUVk5M

URL : https://github.com/nikkula/gitOperationsDemo.git

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
1.  git config --global user.name "xxxx"
2.  git config -global user.email xxxx@xxx.com
3.  git config --list
//pwd -> Present working directory
//After adding new files to the folder then run status command
4.  git status
5.  git add .
6.  git status
7.  git commit -m "Commit Message any text"
8.  git log

//Push the local changes to git Repo
9.  git remote add origin https://github.com/nikkula/gitOperationsDemo.git
10.  git remote -v
11.  git push -u origin master
--------------------------
Step 2:
//Downloading the git remote repository code to my local repository

12.  git pull https://github.com/nikkula/gitOperationsDemo.git

