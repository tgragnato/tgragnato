---
title: Vuln Server
---

```c
#include <stdio.h>
#include <stdlib.h>
#include <winsock.h>

#define IP_ADDRESS  ""
#define PORT_NUMBER 31337
#define BUF_SIZE    2048

void process(SOCKET);

void main(int argc, char *argv[]) {
	SOCKET sock;
	SOCKET new_sock;
	struct sockaddr_in sa;
	WSADATA wsaData;

	if (WSAStartup(MAKEWORD(2, 0), &wsaData) != 0) {
		fprintf(stderr,"WSAStartup() failed");
		exit(1);
	}

	printf("Initializing the server...\n");
	memset(&sa, 0, sizeof(struct sockaddr_in));

	sa.sin_family = AF_INET;
    // bind to a specific address
	//sa.sin_addr.s_addr = inet_addr(IP_ADDRESS);
    // bind to all addresses
	sa.sin_addr.s_addr = htonl(INADDR_ANY);
	sa.sin_port = htons(PORT_NUMBER);
	sock = socket(AF_INET, SOCK_STREAM, 0);

	if (sock < 0) {
		printf("Could not create socket server...\n");
		exit(1);
	}
		
	if (bind(sock, (struct sockaddr *)&sa, sizeof(struct sockaddr_in)) == SOCKET_ERROR) {
		closesocket(sock);
		printf("Could not bind socket...\n");
		exit(1);
	}

	printf("Listening for connections...\n");
	listen(sock, 10);
	while (1) {
		new_sock = accept(sock, NULL, NULL);
    	printf("Client connected\n");

		if (new_sock < 0) {
			printf("Error waiting for new connection!\n");
			exit(1);
		}
		
		process(new_sock);
		closesocket(new_sock);
	}
}

void process(SOCKET sock) {
    char* tmp;
	char buf[1024]; // whoops! should've used the defined BUF_SIZE here
    int i;
    
	memset(&buf, 0, sizeof(buf));
	tmp = malloc(BUF_SIZE);

    i = recv(sock, tmp, BUF_SIZE-1, 0);
    tmp[i+1] = '\x00';
    
    // this would be alright if buf and tmp were the same size
    strcpy(&buf, tmp);

    free(tmp);
	// print data
	printf("Got message:\n%s\n", buf);
}
```

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h> 
#include <sys/socket.h>
#include <netinet/in.h>

#define IP_ADDRESS  "127.0.0.1"
#define PORT_NUMBER 31337
#define BUF_SIZE    2048

void process(int);

int main(int argc, char *argv[]) {
    int sock;
    int new_sock;
    int cli_len;
    struct sockaddr_in sa, ca;

    printf("Initializing the server...\n");
    memset(&sa, 0, sizeof(struct sockaddr_in));
    memset(&ca, 0, sizeof(struct sockaddr_in));

    sa.sin_family = AF_INET;
    // bind to a specific address
    //sa.sin_addr.s_addr = inet_addr(IP_ADDRESS);
    // bind to all addresses
    sa.sin_addr.s_addr = INADDR_ANY;
    sa.sin_port = htons(PORT_NUMBER);
    sock = socket(AF_INET, SOCK_STREAM, 0);

    if (sock < 0) {
        printf("Could not create socket server...\n");
        exit(1);
    }

    if (bind(sock, (struct sockaddr *)&sa, sizeof(struct sockaddr_in)) < 0) {
        close(sock);
        printf("Could not bind socket...\n");
        exit(1);
    }

    printf("Listening for connections...\n");
    listen(sock, 10);
    while (1) {
        cli_len = sizeof(ca);
        new_sock = accept(sock, (struct sockaddr *) &ca, &cli_len);
        printf("Client connected\n");

        if (new_sock < 0) {
            printf("Error waiting for new connection!\n");
            exit(1);
        }

        process(new_sock);
        close(new_sock);
    }

    return 0;
}

void process(int sock) {
    char* tmp;
    char buf[1024]; // whoops! should've used the defined BUF_SIZE here
    int i;

    memset(&buf, 0, sizeof(buf));
    tmp = malloc(BUF_SIZE);

    i = recv(sock, tmp, BUF_SIZE-1, 0);
    tmp[i+1] = '\x00';

    // this would be alright if buf and tmp were the same size
    strcpy(&buf, tmp);

    // ergh... seg fault would be detected here on free by gcc
    //free(tmp);
    // print data
    printf("Got message:\n%s\n", buf);
}
```

```c
#include <stdio.h>
#include <stdlib.h>
#include <winsock.h>

#define IP_ADDRESS     ""
#define PORT_NUMBER    31337
#define PASSWORD       "im_so_lonely\n"
#define BUF_SIZE       128

int auth(SOCKET);

void main(int argc, char *argv[]) {
	SOCKET sock;
	SOCKET new_sock;
	struct sockaddr_in sa;
	WSADATA wsaData;

	if (WSAStartup(MAKEWORD(2, 0), &wsaData) != 0) {
		fprintf(stderr,"WSAStartup() failed");
		exit(1);
	}

	printf("Initializing the server...\n");
	memset(&sa, 0, sizeof(struct sockaddr_in));

	sa.sin_family = AF_INET;
    // bind to a specific address
	//sa.sin_addr.s_addr = inet_addr(IP_ADDRESS);
    // bind to all addresses
	sa.sin_addr.s_addr = htonl(INADDR_ANY);
	sa.sin_port = htons(PORT_NUMBER);
	sock = socket(AF_INET, SOCK_STREAM, 0);

	if (sock < 0) {
		printf("Could not create socket server...\n");
		exit(1);
	}
		
	if (bind(sock, (struct sockaddr *)&sa, sizeof(struct sockaddr_in)) == SOCKET_ERROR) {
		closesocket(sock);
		printf("Could not bind socket...\n");
		exit(1);
	}

	printf("Listening for connections...\n");
	listen(sock, 10);
	while (1) {
		new_sock = accept(sock, NULL, NULL);
    	printf("Client connected\n");

		if (new_sock < 0) {
			printf("Error waiting for new connection!\n");
			exit(1);
		}
		
		if (auth(new_sock)) {
			send(new_sock, "\n***Access granted***\n\n", 23, 0);
			printf("\n***Access granted***\n\n");
		}
		else {
			send(new_sock, "\n***Access denied***\n\n", 22, 0);
			printf("\n***Access denied***\n\n");
		}
		closesocket(new_sock);		
	}
}

int auth(SOCKET sock) {
	char buf[BUF_SIZE];
	int i = 0;
	int res;

	memset(&buf, 0, sizeof(buf));

	// send prompt
	send(sock, "Enter Password: ", 16, 0);

	// read data
	i = recv(sock, buf, BUF_SIZE - 1, 0);
    buf[i] = 0;

	if (strcmp(buf, PASSWORD) == 0)
		res = 1;
	else
		res = 0;
	
	// print data
	printf("Got password: ");
	printf(buf);

	return res;
}
```

```c
#include <stdio.h>
#include <stdlib.h>
#include <winsock.h>

#define IP_ADDRESS     ""
#define PORT_NUMBER    31337
#define BUF_SIZE       1461

void process(SOCKET);

void main(int argc, char *argv[]) {
	SOCKET sock;
	SOCKET new_sock;
	struct sockaddr_in sa;
	WSADATA wsaData;

	if (WSAStartup(MAKEWORD(2, 0), &wsaData) != 0) {
		fprintf(stderr,"WSAStartup() failed");
		exit(1);
	}

	printf("Initializing the server...\n");
	memset(&sa, 0, sizeof(struct sockaddr_in));

	sa.sin_family = AF_INET;
    // bind to a specific address
	//sa.sin_addr.s_addr = inet_addr(IP_ADDRESS);
    // bind to all addresses
	sa.sin_addr.s_addr = htonl(INADDR_ANY);
	sa.sin_port = htons(PORT_NUMBER);
	sock = socket(AF_INET, SOCK_STREAM, 0);

	if (sock < 0) {
		printf("Could not create socket server...\n");
		exit(1);
	}
		
	if (bind(sock, (struct sockaddr *)&sa, sizeof(struct sockaddr_in)) == SOCKET_ERROR) {
		closesocket(sock);
		printf("Could not bind socket...\n");
		exit(1);
	}

	printf("Listening for connections...\n");
	listen(sock, 10);
	while (1) {
		new_sock = accept(sock, NULL, NULL);
    	printf("Client connected\n");

		if (new_sock < 0) {
			printf("Error waiting for new connection!\n");
			exit(1);
		}
		
		process(new_sock);
		closesocket(new_sock);		
	}
}

void process(SOCKET sock) {
    char buf[BUF_SIZE];
    int i = 0;

	memset(&buf, 0, sizeof(buf));

	// read data
	i = recv(sock, buf, BUF_SIZE - 1, 0);
    buf[i] = 0;
	
	// print data
	printf("Got message: ");
	printf(buf);
	printf("\n");
}
```