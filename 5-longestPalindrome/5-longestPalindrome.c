//
// Created by renzheng on 2020/12/21.
//

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char * longestPalindrome(char * s){
    int n = strlen(s);
    int **dp = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 0; i < n + 1; i++) {
        dp[i] = (int *)malloc(sizeof(int) * (n + 1));
        memset(dp[i], 0, sizeof(int) * (n + 1));
    }
    for (int i = 0, j = 0; i < n; i++, j++) {
        dp[i][j] = 1;
    }
    char *ans = (char *)malloc(sizeof(char) * (n + 1));
    strncpy(ans, s, n + 1);
    int maxlen = 0;
    int start = 0;
    int end = 0;
    for (int l = 0; l < n + 1; ++l) {
        for (int i = 0; i + l < n; ++i) {
            int j = i + l;
            if (l == 0) {
                //一个字符是回文字符串
                dp[i][j] = 1;
            } else if (l == 1) {
                //两个字符相等，是回文字符串
                dp[i][j] = (s[i] == s[j]);
            } else {
                dp[i][j] = dp[i+1][j-1] && (s[i] == s[j]);
            }
            if (dp[i][j] && (l + 1 > maxlen)) {
                maxlen = l + 1;
                start = i;
                end = j;
            }
        }
    }
    ans[end+1] = '\0';
    ans += start;

    for (int i = 0; i < n; i++) {
        free(dp[i]);
    }
    free(dp);

    return ans;
}
//gcc -O -g -fsanitize=address main.c -o 5-longestPalindrome
int main() {
    printf("Hello, World!\n");
    char *a = "babad";
    char *b;
    b = longestPalindrome(a);
    printf("%s\n", b);
    free(b);
    return 0;
}

