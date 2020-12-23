//
// Created by renzheng on 2020/12/22.
//

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int lengthOfLongestSubstring(char *s)
{
    //��Ӵ���ʼֵΪ0
    int maxlength = 0;
    int sublength = 0;
    int length = strlen(s);
    if (length == 0 || length == 1) return length;
    char *begin = NULL, *end = NULL, *pos = NULL;
    char *substr = (char *)malloc(sizeof(char) * (length + 1));

    begin = s;
    end = s;

    for (int j = 1; j < length; j++) {
        //�Ӵ�����
        sublength = end - begin + 1;
        maxlength = maxlength < sublength?sublength:maxlength;

        //��s��ȡ�ִ�����substr��
        substr = strncpy(substr, begin, sublength);
        substr[sublength] = '\0';

        if ((pos = strchr(substr, s[j])) != NULL) {
            //�����ظ��ַ���begin �ƶ��� s�ж�Ӧλ�õĺ�һλ
            begin += pos - substr + 1;
        }
        end = s + j;
        sublength = end - begin + 1;
        maxlength = maxlength < sublength?sublength:maxlength;
    }
    free(substr);
    return maxlength;
}

int main()
{
    int maxlength = 0;
    char *s = "a";
    maxlength = lengthOfLongestSubstring(s);
    printf("max length is %d\n", maxlength);
    return 0;
}