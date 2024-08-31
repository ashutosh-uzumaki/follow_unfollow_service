// setup neo4j
//Download docker
// pull the latest image of neo4j -> docker pull neo4j:latest
// run the image -> docker run -name neo4j -p 7474:7474 -p 7687:7687 -d -v $HOME/neo4j/data:/data \ -e NEO4J_AUTH=neo4j/password neo4j:latest

//Now the neo4j database is ready to run.


// RUN THE GOLANG SERVICE
// download the repo
// if versioning error in go.mod -> run go mod vendor
// now directly run go run main.go

QUESTIONS

1. WHY DID I CHOOSE NEO4J?
    NEO4J IS DB BASED ON GRAPH, SO IN FUTURE IF NEEDED WE CAN EXTEND THIS SERVICE TO SUGGEST FOLLOWERS AND OTHER VIDEOS IF ANY OF THE FOLLOWEE WATCHED IN THE FOLLOWING LIST
2. WHY DID I CHOOSE GOLANG?
    THIS IS A SIMPLE VERSION, BUT THERE CAN BE A CASE WHEN CONCURRENT FOLLOW/UNFOLLOW MIGHT BE PERFORMED, HERE GO CONCURRENCY BE UTILIZED WHICH IS EASY TO USE AND ARE LIGHT WEIGHT.

    GO MAINTAINS IT'S GOROUTINE SO WE DON'T HAVE TO WORRY ABOUT MANAGING.