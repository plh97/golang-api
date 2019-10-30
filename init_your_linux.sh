apt update -y &&
apt upgrade -y &&
apt autoremove -y &&
apt install zsh &&
touch ~/.gitconfig &&
echo '[user]\n	name = pengliheng\n	email = pengliheng@gmail.com' >> ~/.gitconfig &&
sh -c "$(wget -O- https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" &&
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions &&
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting &&
wget https://nodejs.org/dist/v12.13.0/node-v12.13.0-linux-x64.tar.xz &&
tar -xf ./node-v12.13.0-linux-x64.tar.xz &&
mv ./node-v12.13.0-linux-x64 /usr/local/ &&
wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz &&
tar -xf ./go1.13.3.linux-amd64.tar.gz &&
mv ./go /usr/local/ &&
cp ./id_rsa_github* /etc/ssh/ &&
vim /etc/ssh/sshd_config &&
service sshd restart