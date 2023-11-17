# Gocron - a API for learning GO language in CLI

## <img src="https://github.com/sinalalebakhsh/Gocron/blob/main/Images/Gocron.jpg" width="100%" align="left">




<br>
	<details>
		<summary>
			Features In Versions
		</summary>
		<details>
			<summary>
			Archive
			</summary>
				<details>
					<summary>
						🎭 V1.0.4 
					</summary>
						🎭 Add Goroutines for searching concurently in data
				</details>
			<details>
				<summary>
					🖌 V1.0.5 
				</summary>
					🖌 Add Single Examples for better learning.
			</details>
			<details>
				<summary>
					🎯 v1.0.6
				</summary>
					🎯 Add Channel and rerlation between Gourotines
			</details>
			<details>
				<summary>
					🧭 v1.0.8
				</summary>
					🧭 Add message for "not found yet"
			</details>
			<details>
				<summary>
					🌍 v1.0.12
				</summary>
					🌍 Adding => All Creating HTTP Servers
			</details>
			<details>
				<summary>
					🍪 v1.0.13
				</summary>
					🍪 Add in everything Inspecting the Request
			</details>
			<details>
				<summary>
					🥎 v1.0.14
				</summary>
					🥎🥎🥎 Creating My First Server
			</details>
			<details>
				<summary>
					🥝 v1.0.15
				</summary>
					🥝 Add example for saving logs in TXT file from Server GO
			</details>
			<details>
				<summary>
					🏆 v1.0.16
				</summary>
					🏆 Will write in Terminal and File for feture analyzing. 
			</details>
			<details>
				<summary>
					🔰 v1.0.17
				</summary>
					🔰 Will write in Terminal and File for feture analyzing. 
			</details>
			<details>
				<summary>
					💎 v1.0.20
				</summary>
					💎 Add this section:
					The user is asked if he wants to enter the project page. 
			</details>
			Complete HTTP Servers First Step
			<details>
				<summary>
					🪙 v1.0.21
				</summary>
					🪙 Complete HTTP Servers First Step 
			</details>
			<details>
				<summary>
					⚽ 1.0.22
				</summary>
					⚽ Add Clear Screen in program After Ending for in current running program you can use it ⚽
			</details>
			<details>
				<summary>
					⏺ v1.0.23
				</summary>
					⏺ 1- add build program
					⏺ 2- debuging
					⏺ 3- add map instead slice"
			</details>
			<details>
				<summary>
					✅ v1.0.24
				</summary>
					✅ develop searching in examples with more than one or two arguments:
					✅  this is better for searching like this:
					✅      http servers example
					✅  result is true for this.
			</details>
			<details>
				<summary>
					🔴 v1.0.25
				</summary>
					🔴 Add two file .sh for pushing in GitHub
			</details> 
			<details>
				<summary>
					🔔 v1.0.25
				</summary>
					🔔 Debug single definition for searching by To Lower Function
						Clear Key Implementation:
							Identify the key that will trigger the clearing functionality. This could be a specific keystroke or command.
							Write the code that listens for this key input and executes the clear action.
							Ensure that the user is provided with feedback confirming that the data has been cleared.
						Enhanced Search Functionality:
							Determine what specific improvements you want to make in the search feature. This could include adding filters (e.g., by category, difficulty level), allowing more complex search queries, or implementing a fuzzy search algorithm.
							Implement the necessary code to realize these improvements. This might involve changes to the user interface and underlying data structures.
			</details> 
			<details>
				<summary>
					🍅 v1.0.27
				</summary>
					🍅 Add license MIT
			</details> 
			<details>
				<summary>
					🌈 v1.0.28
				</summary>
					🌈 Add a feature for Question/Answer from user
			</details> 
			<details>
				<summary>
					🟣 v1.0.29
				</summary>
					🟣 Add a feature if user wrote 'godoc usage' specified
			</details> 
			<details>
				<summary>
					🐻 v1.0.31
				</summary>
					🐻 Debuging For being a GitHub donkey!!!!!
			</details>
			<details>
				<summary>
					💠 v1.0.32
				</summary>
					💠💠💠 If User Input Is Just Number 💠💠💠
			</details>
			<details>
					<summary>
						💠 v1.0.33
					</summary>
					💠💠💠 💠💠💠 Adding There are complete => Numbers for Keys Map until 480 💠💠💠
			</details>
		</details>
		<details>
				<summary>
					v1.0.38
				</summary>
			Now for Windows and Ubuntu amd64 = Building Executables for Different Architectures
			This is my First developed project , so maybe will not work for your Windows
			I learn this from:
			https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04
		</details>
	</details>
</br>


# Indtroduction 
is a API for learning GO language with example.
my name is Sina LalehBakhsh, I hope this API is useful for you
after running program, write your single word about any of GO language.
if your perpuse is more than one word, for convenience searching, just write keywords.
like this:
map slice


## 1.Download Gocron
	git clone https://github.com/sinalalebakhsh/Gocron.git

## 2.in Terminal write current directory dowloaded:
	./Gocron

## 3.Use it
	just this!



## GO installation
### The first way
---------------------------------------
#### 1-1 download from 
	https://go.dev/dl/ 

#### 1-2 Linux:
	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
	export PATH=$PATH:/usr/local/go/bin

#### 1-3 Verify that in Terminal:
	go version

### Debug Installation:
#### Remove any previous Go installation
	sudo rm -rf /usr/local/go 
	sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
	echo $PATH | grep "/usr/local/go/bin"

#### for config to all directories:
	echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
	echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.bashrc
	echo 'export GOPATH=$HOME/go' >> ~/.bashrc
	source ~/.profile


##### Or use this for only you want to use Go specific directory:
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH


### 2-Using
-----------------------------------------

You can use this command before start:

#### ◉ Run
	go run .
	
#### ◉ Build
	go build .
	than:
	./Gocron
	
#### ◉ Help
	./Gocron -h 
	./Gocron help
	./Gocron -help
	./Gocron --help

#### ◉ Show All
	./Gocron all
	./Gocron -all
	./Gocron --all


### ◉ if just write one input you get all about that:
	READING AND WRITING DATA
    READINGANDWRITINGDATA
    ALL READING AND WRITING DATA
    ALLREADINGANDWRITINGDATA
    ALL REGEX 
    ALLREGEX
    ALL TIME
    ALLTIME
    ALL HTML AND TEMPLATE
	ALLHTMLANDTEMPLATE
    ALL WORKING WITH FILES
	ALLWORKINGWITHFILES
    ALL JSON
	ALLJSON
	ALL JSON DATA
	ALLJSONDATA
	ALL WORK WITH JSON DATA
	ALLWORKWITHJSONDATA
	ALL WORKING WITH JSON DATA
	ALLWORKINGWITHJSONDATA
	ALL HTTP SERVERS
    ALLHTTPSERVERS
    ALL CREATING HTTP SERVERS
    ALLCREATINGHTTPSERVERS
	ALL HTTP CLIENTS
	ALLHTTPCLIENTS
	ALL CREATING HTTP CLIENTS
	ALLCREATINGHTTPCLIENTS
	ALL WORKING WITH DATABASES
	ALLWORKINGWITHDATABASES
	ALL DATABASES
	ALLDATABASES
	ALLUSINGREFLECTION
	ALL USING REFLECTION

### ◉ 💠 v1.0.32
💠💠💠 If User Input Is Just Number 💠💠💠
### example after run Gocron:
	operator 1 question
### result is this:
	The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.	Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, Note: the return value should be a float64, from 0 to 100:
another example after run Gocron:

🌎🌏🌐🏀⚾🥎🏐🟠🟡🟢🔵🟤🔘🍘🥘🫓🏵🍑🍊🍈🌳❤️🎾🏕🏜🍺🫖🎳🎖🥇🥈🥉🎃🔥💧⭐🌟🏎🏍🚂🟥🟧🟨🟩🟦🟪🟫🔶🔷💠🔆⏹🛡🔒📌🗂📂💰📔📕📖📗📘📙📚📓📒💍🎨
