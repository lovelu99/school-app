# school app

cd frontend/
npm install
npm start


cd employeeservice/

python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
export MONGODB_URI="mongodb://admin:password123@172.18.0.2:27017/schooldb?authSource=admin"
export DATABASE_NAME="schooldb"
echo $MONGODB_URI
echo $DATABASE_NAME
python app.py


cd studentservice/

wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
Should show:
go version go1.23.4 linux/amd64

export MONGODB_URI="mongodb://admin:password123@172.18.0.2:27017/schooldb?authSource=admin"
export DATABASE_NAME="schooldb"
go mod tidy
go run main.go


cd teacherservice/
sudo apt update
sudo apt install -y openjdk-17-jre
sudo apt install -y  maven

mvn clean package
export MONGODB_URI="mongodb://admin:password123@172.18.0.2:27017/schooldb?authSource=admin"
export DATABASE_NAME="schooldb"
java -jar target/teacherservice-1.0.0.jar
