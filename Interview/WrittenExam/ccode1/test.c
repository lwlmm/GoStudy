#include<stdio.h>
#include<string.h>

int main() {
    int i, j, n, m, cnt, cntLine, cntPool;
    scanf("%d\n", &n);
    char line[200];
    char pool[151][11];
    char lineC[15][11];
    char word[11];
    for (i = 0; i < n; i++) {
        strcpy(line, "");
        strcpy(word, "");
        for (j = 0; j < 15; j++) {
            strcpy(lineC[j], "");
        }
        gets(&line);
        cnt = 0;
        cntLine = 0;
        for (j = 0; j < strlen(line)-1; j++) {
            if (line[j] != ' ') {
                word[cnt] = line[j];
                cnt++;
            } else {
                cnt = 0;
                for (j = 0; j < 15; j++) {
                    if (strcmp(lineC[j], word) == 0) {
                        break;
                    }
                }
                strcpy(lineC[cntLine], word);
                strcpy(word, "");
                cntLine++;
            }
        }
        for (j = 0; j < 15; j++) {
            if (strcmp(lineC[j], "") != 0) {
                strcpy(pool[cntPool], lineC[j]);
                cntPool++;
            }
        }
    }
    scanf("%d\n", &m);
    for (i = 0; i < m; i++) {
        strcpy(word, "");
        scanf("%s\n", &word);
        cnt = 0;
        for (j = 0; j < 151; j++) {
            if (strcmp(pool[j], word) == 0) {
                cnt++;
            }
        }
        printf("%d\n", cnt);
    }
    return 0;
}