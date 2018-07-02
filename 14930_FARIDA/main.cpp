#include<bits/stdc++.h>

using namespace std;

int main(int argc, char** argv) {

    long long int T, N, x;
    scanf("%lld", &T);

    for (long long int t = 0; t < T; t++) {
        scanf("%lld", &N);
        if (N == 0) {
            cin.ignore();
            printf("Case %lld: 0\n", t+1);
            continue;
        }

        long long int inner = 0, outer = 0, _inner = 0;

        scanf("%lld", &x);
        inner = x;

        for (long long int n = 0; n < N - 1; n++) {
            scanf("%lld", &x);
            _inner = inner;
            inner = outer + x;
            outer = max(outer, _inner);
        }

        printf("Case %lld: %lld\n", t+1, max(outer, inner));

    }


    return 0;
}