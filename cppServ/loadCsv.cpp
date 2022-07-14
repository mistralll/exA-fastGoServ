#include "setting.h"

void ResetTags() {
    rep(i, sizeof(tags)) {
        rep(j, sizeof(tags[i].img)) {
            tags[i].img[j] = -1;
        }
    }
    cout << "vec: tags is reset." << endl;
}

vector<string> split(string &input, char delimiter) {
    istringstream stream(input);
    string field;
    vector<string> result;
    while(getline(stream, field, delimiter)) {
        result.push_back(field);
    }
    return result;
}

void LoadTagCsv(string filename) {
    ResetTags();

    cout << "now loading tag file..." << endl;

    ifstream ifs(filename);
    if (!ifs) {
        cout << "Fail to open tag.csv file." << endl;
    }
    
    ll tagcnt = 0; ll imgcnt = 0;

    string line;
    while(getline(ifs, line)) {
        vector<string> strvec = split(line, ',');

        if(strvec[0] == "") {
            tags[tagcnt].img[imgcnt] = line[1];
            imgcnt++;
        } else {
            if(tagcnt != 0) tagcnt++;
            imgcnt = 0;
        }
    }

    cout << "complited!" << endl;
}