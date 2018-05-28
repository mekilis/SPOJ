#include <cstdio>
#include <bits/stdc++.h>
using namespace std;

void toInt(int &num) {
    bool negative = false;
    char c;

    num = 0;

    c = getchar();
    
    while (c == EOF) {
		c = getchar();
	}
	
    if (c == '-') {
        negative = true;
        c = getchar();
    }

    for (; (c > 47 && c < 58); c = getchar())
        num = 10 * num + c - 48;
    
    if (negative)
        num *= -1;
}

struct Node {
    int maxPrefixSum;
    int maxSuffixSum;
    int totalSum;
    int maxSubarraySum;
 
    Node()
    {
        maxPrefixSum = maxSuffixSum = maxSubarraySum = -15008;
        totalSum = 0;
    }
};

Node merge(Node leftChild, Node rightChild)
{
    Node parentNode;
    parentNode.maxPrefixSum = max(leftChild.maxPrefixSum, leftChild.totalSum + rightChild.maxPrefixSum);
 
    parentNode.maxSuffixSum = max(rightChild.maxSuffixSum, rightChild.totalSum + leftChild.maxSuffixSum);
 
    parentNode.totalSum = leftChild.totalSum + rightChild.totalSum;
 
    parentNode.maxSubarraySum = max(max(leftChild.maxSubarraySum, rightChild.maxSubarraySum),
        leftChild.maxSuffixSum + rightChild.maxPrefixSum);
 
    return parentNode;
}

void constructTree(Node* tree, int arr[], int start, int end, int index)
{
    if (start == end) {
        tree[index].totalSum = arr[start];
        tree[index].maxSuffixSum = arr[start];
        tree[index].maxPrefixSum = arr[start];
        tree[index].maxSubarraySum = arr[start];

        return;
    }
 
    int mid = (start + end) / 2;
    constructTree(tree, arr, start, mid, 2 * index);
    constructTree(tree, arr, mid + 1, end, 2 * index + 1);

    tree[index] = merge(tree[2 * index], tree[2 * index + 1]);
}

Node query(Node* tree, int ss, int se, int qs, int qe, int index)
{
    if (ss > qe || se < qs) {
        Node n;
        return n;
    }

    if (ss >= qs && se <= qe) {
        return tree[index];
    }
 
    int mid = (ss + se) / 2;
    Node left = query(tree, ss, mid, qs, qe, 2 * index);
    Node right = query(tree, mid + 1, se, qs, qe, 2 * index + 1);
 
    Node res = merge(left, right);
    return res;
}

int *A;

int main() {
    // ios_base::sync_stdio(false);
    // cin.tle(NULL);
    int N, M, n, m, x, i, j;
    scanf("%d", &N);
    A = new int[N+1];

    for (n = 0; n < N; n++) {
        scanf("%d", &x);
        A[n] = x;
    }

    x = (int)(ceil(log2(N)));
    int max_size = 2 * (int)pow(2, x) - 1; 
    Node* tree = new Node[max_size];
 
    constructTree(tree, A, 0, N - 1, 1);

    scanf("%d", &M);

    for (m = 0; m < M; m++) {
        scanf("%d", &i);
        scanf("%d", &j);

        Node node = query(tree, 0, N-1, i-1, j-1, 1);

        printf("%d\n", node.maxSubarraySum);
    }

    delete [] A;
    return 0;
}

// solved in 0.16s 18M used.
