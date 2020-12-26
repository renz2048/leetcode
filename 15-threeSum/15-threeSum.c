//
// Created by renzheng on 2020/12/24.
//

#include <stdio.h>
#include <stdlib.h>

#define DEFAULT_SIZE 10000

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** threeSum(int* nums, int numsSize, int* returnSize, int** returnColumnSizes)
{
    returnColumnSizes = (int **)malloc(sizeof(int *) * DEFAULT_SIZE);
    int *a = NULL, *b = NULL, *c = NULL;
    *returnSize = 0;
    if (numsSize < 3) {
        return returnColumnSizes;
    }
    for (int i = 0; i < numsSize - 2; ++i) {
        a = nums + i;
        for (int j = i + 1; j < numsSize - 1; ++j) {
            b = nums + j;
            for (int k = j + 1; k < numsSize; ++k) {
                c = nums + k;
                if (*a + *b + *c == 0) {
                    if (!((*returnSize) % DEFAULT_SIZE) && (*returnSize) >= DEFAULT_SIZE) {
                        returnColumnSizes = (int **)realloc(returnColumnSizes, *returnSize + DEFAULT_SIZE);
                    }
                    returnColumnSizes[*returnSize] = (int *)malloc(sizeof(int) * 4);
                    returnColumnSizes[*returnSize][0] = *a;
                    returnColumnSizes[*returnSize][1] = *b;
                    returnColumnSizes[*returnSize][2] = *c;
                    (*returnSize)++;
                }
            }
        }
    }
    return returnColumnSizes;
}

int main()
{
    int **returnColumnSizes;
    int *returnSize = (int *)malloc(sizeof(int));
    //int nums[6] = {-1, 0, 1, 2, -1, -4};
    int nums[6] = {-1, 0, 1, 2};
    int numSize = 4;
    returnColumnSizes = threeSum(nums, numSize, returnSize, returnColumnSizes);

    for (int i = 0; i < *returnSize; ++i) {
        for (int j = 0; j < 3; ++j) {
            printf("%d ", returnColumnSizes[i][j]);
        }
        printf("\n");
    }
    free(returnColumnSizes);
    return 0;
}