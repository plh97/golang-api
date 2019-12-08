apt update -y &&
apt upgrade -y &&
apt autoremove -y &&
touch ~/.gitconfig &&
echo '[user]' >> ~/.gitconfig &&
echo '	name = pengliheng' >> ~/.gitconfig &&
echo '	email = pengliheng111@gmail.com' >> ~/.gitconfig &&
wget https://nodejs.org/dist/v12.13.0/node-v12.13.0-linux-x64.tar.xz &&
tar -xf ./node-v12.13.0-linux-x64.tar.xz &&
mv ./node-v12.13.0-linux-x64 /usr/local/node &&
wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz &&
tar -xf ./go1.13.3.linux-amd64.tar.gz &&
mv ./go /usr/local/go &&
rm ./node-v12.13.0-linux-x64.tar.xz &&
rm ./go1.13.3.linux-amd64.tar.gz &&
cp ./id_rsa_github* /etc/ssh/ &&
git config --global credential.helper store &&
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common -y &&
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - &&
wget "https://download.docker.com/linux/ubuntu/dists/disco/pool/stable/amd64/containerd.io_1.2.6-3_amd64.deb" &&
wget "https://download.docker.com/linux/ubuntu/dists/disco/pool/stable/amd64/docker-ce-cli_19.03.3~3-0~ubuntu-disco_amd64.deb" &&
wget "https://download.docker.com/linux/ubuntu/dists/disco/pool/stable/amd64/docker-ce_19.03.3~3-0~ubuntu-disco_amd64.deb" &&
sudo dpkg -i "containerd.io_1.2.6-3_amd64.deb" &&
sudo dpkg -i "docker-ce-cli_19.03.3~3-0~ubuntu-disco_amd64.deb" &&
sudo dpkg -i "docker-ce_19.03.3~3-0~ubuntu-disco_amd64.deb" &&sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
docker version &&
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose &&
sudo chmod +x /usr/local/bin/docker-compose &&
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose &&
docker-compose --version &&
rm containerd.io_1.2.6-3_amd64.deb &&
rm docker-ce_19.03.3~3-0~ubuntu-disco_amd64.deb &&
rm docker-ce-cli_19.03.3~3-0~ubuntu-disco_amd64.deb &&
apt-get install build-essential -y
# vim /etc/ssh/sshd_config