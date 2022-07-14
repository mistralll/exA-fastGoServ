#include "setting.h"

void test_printTags() {
    rep(i, sizeof(tags)) {
        cout << tags[i].name << endl;
        rep(j, sizeof(tags[i].img)) {
            if (tags[i].img[j] != -1) cout << tags[i].img[j] << ", ";
        }
        cout << endl;

        break;
    }
}