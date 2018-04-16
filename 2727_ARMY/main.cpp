#include <iostream>
#include <stdio.h>

using namespace std;

int main() {
    int T, NG, NM;
    int NGMax = -1;
    int NMMax = -1;
    scanf("%d",&T);

    for (int t = 0; t < T; t++) {
        scanf("%d",&NG);
        scanf("%d",&NM);

        int NGArr[NG];
        int NMArr[NM];

        for (int ng = 0; ng < NG; ng++) {
            scanf("%d",&NGArr[ng]);
        }

        for (int nm = 0; nm < NM; nm++) {
            scanf("%d",&NMArr[nm]);
        }

        NGMax = NGArr[0];
        NMMax = NMArr[0];

        for (int i = 1; i < NG; i++) {
            if (NGMax < NGArr[i]) {
                NGMax = NGArr[i];
            }
        }

        for (int i = 1; i < NM; i++) {
            if (NMMax < NMArr[i]) {
                NMMax = NMArr[i];
            }
        }

        if (NGMax >= NMMax) {
            printf("Godzilla\n");
        } else if (NGMax < NMMax) {
            printf("MechaGodzilla\n");
        } else {
            printf("uncertain\n");
        }
    }

    return 0;
}
