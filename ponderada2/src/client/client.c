#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>

#define SERVER_PORT 12345  
#define BUF_SIZE 4096      

void fatal(char *string) {
    printf("%s\n", string);
    exit(1);
}

int main(int argc, char **argv) {
    int c, s, bytes;
    char buf[BUF_SIZE]; 
    struct hostent *h;   
    struct sockaddr_in channel; /

    if (argc != 3) fatal("Usage: client server-name file-name");
    h = gethostbyname(argv[1]);
    if (!h) fatal("gethostbyname failed");
    
    s = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (s < 0) fatal("socket");
    
    memset(&channel, 0, sizeof(channel));
    channel.sin_family = AF_INET;
    memcpy(&channel.sin_addr.s_addr, h->h_addr, h->h_length);
    channel.sin_port = htons(SERVER_PORT);
    
    c = connect(s, (struct sockaddr *)&channel, sizeof(channel));
    if (c < 0) fatal("connect failed");
    
    write(s, argv[2], strlen(argv[2]) + 1);
    
    while (1) {
        bytes = read(s, buf, BUF_SIZE);
        if (bytes <= 0) exit(0);
        write(1, buf, bytes);
    }
}
