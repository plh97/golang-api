apt update -y &&
apt upgrade -y &&
apt autoremove -y &&
touch ~/.gitconfig &&
echo '[user]' >> ~/.gitconfig &&
echo '	name = pengliheng' >> ~/.gitconfig &&
echo '	email = pengliheng@gmail.com' >> ~/.gitconfig &&
wget https://nodejs.org/dist/v12.13.0/node-v12.13.0-linux-x64.tar.xz &&
tar -xf ./node-v12.13.0-linux-x64.tar.xz &&
mv ./node-v12.13.0-linux-x64 /usr/local/ &&
wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz &&
tar -xf ./go1.13.3.linux-amd64.tar.gz &&
mv ./go /usr/local/ &&
rm ./node-v12.13.0-linux-x64.tar.xz &&
rm ./go1.13.3.linux-amd64.tar.gz &&
cp ./id_rsa_github* /etc/ssh/ &&
git config --global credential.helper store &&
vim /etc/ssh/sshd_config