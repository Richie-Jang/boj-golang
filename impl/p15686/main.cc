#include <stdio.h>

const int N = 5;
const int M = 3;
int data[N] = {1,2,3,4,5};
int selData[M];
int perMuteCount = 0;

void printCombination(int curSel, int curPos) {
    if (curSel == M) {
        // print
        for (auto k : selData) {
            printf("%d ", k);
        }
        printf("\n");
        return;
    }

    for (int i = curPos; i < N; i++) {
        selData[curSel] = data[i];
        printCombination(curSel+1, i+1);
    }
}

void swap(int st, int et) {
    int t = selData[st];
    selData[st] = selData[et];
    selData[et] = t;
}

void permute (int start, int end) {
    if (start == end) {
        for (const auto i : selData) {
            printf("%d ", i);
        }
        printf("\n");
        perMuteCount++;
    } else {
        for (int i = start; i <= end; i++) {
            swap(start, i);
            permute(start+1, end);
            swap(start, i);
        }
    }
}

void printPermutation(int curSel, int curPos) {
    if (curSel == M) {
        permute(0, M-1);
        return;
    }

    for (int i = curPos; i < N; i++) {
        selData[curSel] = data[i];
        printPermutation(curSel+1, i+1);
    }
}

int main() {

    printPermutation(0, 0);
    printf("Total: %d\n", perMuteCount);
    return 0;
}
