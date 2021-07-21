#include<stdio.h>
#include<string.h>

int main() {
    int i, j, n, m, cnt;
    scanf("%d\n", &n);
    char pool[151][11];
    char word[11];
    for (i = 0; ; i++) {
        strcpy(word, "");
        scanf("%s\n", &word);
        if (word[0] >= '0' && word[0] <='9') {
            break;
        }
        strcpy(pool[cnt], word);
        cnt++;
    }
    m = atoi(word);
    for (i = 0; i < m; i++) {
        strcpy(word, "");
        scanf("%s\n", &word);
        cnt = 0;
        if (strlen(word) > 1) {
            printf ("1\n");
            continue;
        }
        for (j = 0; j < 151; j++) {
            if (strcmp(pool[j], word) == 0) {
                cnt++;
            }
        }
        printf("%d\n", cnt);
    }
    return 0;
}