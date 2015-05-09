#ifndef _COMMUNICATE__
#define _COMMUNICATE_

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h> 


using namespace std;

class Communicate{
public:
  Communicate();
  void init();

  void send(string msg);

  void end();
private:
  int sockfd, portno, n, buffbytes;
  struct sockaddr_in serv_addr;
  struct hostent *server;
  void error(const char *msg);
  char buffer[256];
};

#endif

