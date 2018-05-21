#include <iostream>
using namespace std;

int main()
{
	int T, h, w, m;
	int **tiles;
	int **tilesCopy;

	cin >> T;
	for (int t = 0; t < T; t++) {
		cin >> h;
		cin >> w;

		tiles = new int*[h];
		tilesCopy = new int*[h];

		for (int i = 0; i < h; i++) {

			tiles[i] = new int[w];
			tilesCopy[i] = new int[w];

			for (int j = 0; j < w; j++) {
				cin >> m;
				tiles[i][j] = m;
				tilesCopy[i][j] = m;
			}
		}

		int max = 0, left, right, down;

		for (int i = 0; i < h-1; i++) {
			for (int j = 0; j < w; j++) {

				down = tilesCopy[i][j] + tilesCopy[i+1][j];
				if (down > tiles[i+1][j])
					tiles[i+1][j] = down;

				if (j > 0 && j < w-1) {
					left = tilesCopy[i][j] + tilesCopy[i+1][j-1];
					right = tilesCopy[i][j] + tilesCopy[i+1][j+1];

					if (left > tiles[i+1][j-1])
						tiles[i+1][j-1] = left;

					if (right > tiles[i+1][j+1])
						tiles[i+1][j+1] = right;

				} else if (j == 0) {
					right = tilesCopy[i][j] + tilesCopy[i+1][j+1];

					if (right > tiles[i+1][j+1])
						tiles[i+1][j+1] = right;
				} else {
					left = tilesCopy[i][j] + tilesCopy[i+1][j-1];

					if (left > tiles[i+1][j-1])
						tiles[i+1][j-1] = left;
				}
			}

			// update row
			tilesCopy[i+1] = tiles[i+1];
			//for (int j = 0; j < w; j++)
			//	tilesCopy[i+1][j] = tiles[i+1][j];
		}

		for (int i = 0; i < w; i++) {
			if (tiles[h-1][i] > max)
				max = tiles[h-1][i];
		}

		cout << max << endl; // no endl caused 1 WA
	}

	return 0;
}

