#!/bin/sh

currDate=$(date)
mkdir "$1 at $currDate"
cd "$1 at $currDate"
mkdir about_me
cd about_me
mkdir personal
mkdir professional
cd personal
touch facebook.txt
echo "https://www.facebook.com/" > facebook.txt
cd ..
cd professional
touch linkedin.txt
echo "https://id.linkedin.com/" > linkedin.txt
cd .. 
cd ..
mkdir my_friends
cd my_friends
touch list_of_my_friends.txt
echo $(curl -XGET 'https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt') >> list_of_my_friends.txt
cd ..
mkdir my_system_info
cd my_system_info
touch about_this_laptop.txt
echo "My username: tegarimansyah" > about_this_laptop.txt
echo "with host: $(uname -a)" >> about_this_laptop.txt
touch internet_connection.txt
echo $(ping google.com) > internet_connection