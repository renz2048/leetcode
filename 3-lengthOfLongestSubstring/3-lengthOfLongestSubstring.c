//
// Created by renzheng on 2020/12/22.
//

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int lengthOfLongestSubstring(char *s)
{
    int maxlength = 0;
    int length = strlen(s);
    char *begin = NULL, *end = NULL, *pos = NULL;
    char *substr = (char *)malloc(sizeof(char) * (length + 1));

    begin = s;
    end = s;
    maxlength = end - begin + 1;

    for (int j = 1; j < length; j++) {
        substr = strncpy(substr, begin, end - begin + 1);
        substr[end - begin + 1 ] = '\0';
        if ((pos = strchr(substr, s[j])) != NULL) {
            //´æÔÚÖØ¸´×Ö·û
            begin += pos - substr;
        }
        maxlength = maxlength < (end - begin + 1)?(end - begin + 1):maxlength;
        end = s + j;
    }
    return maxlength;
}

int main()
{
    int maxlength = 0;
    char *s = "pwwkew";
    maxlength = lengthOfLongestSubstring(s);
    printf("max length is %d\n", maxlength);
    return 0;
}