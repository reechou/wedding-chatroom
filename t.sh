# maybe more powerful
# for mac (sed for linux is different)
dir=`echo ${PWD##*/}`
grep "wedding-chatroom" * -R | grep -v Godeps | awk -F: '{print $1}' | sort | uniq | xargs sed -i '' "s#wedding-chatroom#$dir#g"
mv wedding-chatroom.ini $dir.ini

