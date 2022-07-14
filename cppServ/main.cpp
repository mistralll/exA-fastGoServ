#include "setting.h"

#include "loadCsv.cpp"
#include "test_print.cpp"

vector<Image> images(10397271);
vector<Tag> tags(860621);


int main() {
    LoadTagCsv("../csv/comp/tagComped.csv");
    test_printTags();
}