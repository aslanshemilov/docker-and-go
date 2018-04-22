Last login: Thu Mar 22 13:24:38 on ttys000
Shilpas-MacBook-Pro:~ shilpakancharla$ cd images
Shilpas-MacBook-Pro:images shilpakancharla$ touch docker-compose.yml
Shilpas-MacBook-Pro:images shilpakancharla$ docker swarm init
Swarm initialized: current node (tdseetdlppy4ptiu2rutvbgvs) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-21z9rya5ll5whuukorwj7f2eqvicryyngmzr70dhat6rual65f-dr3kdlc9ptagsirtyx8wlw6oh 192.168.65.3:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.

Shilpas-MacBook-Pro:images shilpakancharla$ docker stack deploy -c docker-compose.yml getstartedlab
Creating network getstartedlab_webnet
Creating service getstartedlab_web
Shilpas-MacBook-Pro:images shilpakancharla$ docker service ls
ID                  NAME                MODE                REPLICAS            IMAGE                    PORTS
6w2h84adki5f        getstartedlab_web   replicated          5/5                 skancharla/repo1:part1   *:80->80/tcp
Shilpas-MacBook-Pro:images shilpakancharla$ docker service ps getstartedlab_web
ID                  NAME                  IMAGE                    NODE                    DESIRED STATE       CURRENT STATE                ERROR               PORTS
phewymq26m9g        getstartedlab_web.1   skancharla/repo1:part1   linuxkit-025000000001   Running             Running about a minute ago                       
t12tsulo1bdl        getstartedlab_web.2   skancharla/repo1:part1   linuxkit-025000000001   Running             Running about a minute ago                       
w81v1pz7bizc        getstartedlab_web.3   skancharla/repo1:part1   linuxkit-025000000001   Running             Running about a minute ago                       
lpucm8a0fabo        getstartedlab_web.4   skancharla/repo1:part1   linuxkit-025000000001   Running             Running about a minute ago                       
nximca96nu0b        getstartedlab_web.5   skancharla/repo1:part1   linuxkit-025000000001   Running             Running about a minute ago                       
Shilpas-MacBook-Pro:images shilpakancharla$ docker container ls -q
b8ad24b05427
2665c68b3db2
4d8defe5e802
aea3e034c786
e7bc23932675
Shilpas-MacBook-Pro:images shilpakancharla$ curl -4 http://localhost
<h3>Hello World!</h3><b>Hostname:</b> e7bc23932675<br/><b>Visits:</b> <i>cannot connect to Redis, counShSShSSSShilpas-MacBook-Pro:images shilpakancharla$ docker stack deploy -c docker-compose.yml getstartedlab
Updating service getstartedlab_web (id: 6w2h84adki5fzj7161wxegfnv)
Shilpas-MacBook-Pro:images shilpakancharla$ docker stack rm getstartedlab
Removing service getstartedlab_web
Removing network getstartedlab_webnet
Shilpas-MacBook-Pro:images shilpakancharla$ docker swarm leave --force
Node left the swarm.
Shilpas-MacBook-Pro:images shilpakancharla$ 
