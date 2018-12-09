# Hadoop: Hello World

Hadoop is a set of tools of big-data storage (HDFS), computing (MapReduce) and resource management (Yarn)
Well, the official document is long and quite user unfriendly. In order to deep dive it quicker, docker is the way to go.

## Install

The easiest way to run Hadoop right away on my Mac of course is using Docker. Assuming you have docker installed.

Here is a really good place [repo https://github.com/HariSekhon/Dockerfiles](https://github.com/HariSekhon/Dockerfiles) for many big-data Dockerfiles. 

1. Clone it 

```shell
$ git clone git@github.com:HariSekhon/Dockerfiles.git
```

2. `cd` to `hadoop-dev` (I assume it's for development purpose) folder and run
```shell
$ docker-compose up
```

3. Checking the available ports... Emm.. so many of them.
```shell
$ docker ps

0.0.0.0:8042->8042/tcp,
8020/tcp,
9000/tcp,
0.0.0.0:8088->8088/tcp,
0.0.0.0:19888->19888/tcp,
0.0.0.0:50010->50010/tcp,
0.0.0.0:50020->50020/tcp,
0.0.0.0:50070->50070/tcp,
0.0.0.0:50075->50075/tcp,
10020/tcp,
0.0.0.0:50090->50090/tcp
```

4. Let's try to open them one by one from browser,

- 8088: the cluster dashboard. There are status about Nodes, Applications, Conf, Scheduler.
- 8042: the a similar dashboard but for single node.
- 8020: Not working
- 9000: Not working
- 50010: Not working
- 50020: > It looks like you are making an HTTP request to a Hadoop IPC port. This is not the correct port for the web interface on this daemon.
- 50070: A much better looking UI for Datanodes, Snapshot, Startup Progress, Logs and even a File System Browser like `Finder`. This looks like the UI for HDFS.
- 50075: A similar UI as 50070 for single Datanode.

5. After peeking the dashboards. I didn't see the UI for adding data. Let's see how should we do it.

[File System Shell](https://hadoop.apache.org/docs/r3.0.1/hadoop-project-dist/hadoop-common/FileSystemShell.html) listed many commands to interact with HDFS.

Attach the command line to the running container and run some Hadoop commands
```shell
$ docker exec -it hadoop-dev_hadoop-dev_1 /bin/bash

$ hadoop fs -mkdir /helloworld
$ hadoop fs -ls /
```

check [File Browser](http://locahost:50070/explorer.html), the `helloworld` dir we just created  is showing in the list. Cool.

Let me put some data into it. I got some data set [link](https://ed-public-download.app.cloud.gov/downloads/Most-Recent-Cohorts-Scorecard-Elements.csv) from [this repo](https://github.com/awesomedata/awesome-public-datasets).

In the container

```shell
$ curl https://ed-public-download.app.cloud.gov/downloads/Most-Recent-Cohorts-Scorecard-Elements.csv > Most-Recent-Cohorts-Scorecard-Elements.csv

$ hadoop fs -put Most-Recent-Cohorts-Scorecard-Elements.csv /helloworld/input
```

And now I can see and download this file in from File Browser too.

6. Let's do some MapReduce with my data now.

Following the [MapReduce Tutorial](https://hadoop.apache.org/docs/r3.0.1/hadoop-mapreduce-client/hadoop-mapreduce-client-core/MapReduceTutorial.html) here, the WordCount MapReducer.

Created the `WordCount.java` file. Trying to compile it by following the steps in the doc but, well, the docker file I used doesn't have `${JAVA_HOME}/lib/tools.jar`. 

An alternative way is

```shell
// returns me the hadoop libraries path
$ hadoop classpath 

// to set it into the HADOOP_CLASSPATH
$ export HADOOP_CLASSPATH=$(hadoop classpath)

// will build to this folder
$ mkdir build 

// compile the java file into build folder
$ javac -cp $HADOOP_CLASSPATH WordCount.java -d build 

// created the JAR file from the built application
$ jar -cvf wc.jar -C build/ .

```

We created the WordCount jar file, now we cn run the WordCount MapReducer against our input data and output it.

```shell
$ hadoop jar wc.jar WordCount /helloworld/input /helloworld/output
```

The job started running and finished with the counts. Coo!

We are done for the first step on Hadoop: installed it and ran the hello world. 

