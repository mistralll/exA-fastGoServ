#ifndef HEADER_H
#define HEADER_H
#pragma once

#include <bits/stdc++.h>
#include <string.h>
#include <fstream>
#include <sstream>

using namespace std;
#define rep(i, n) for (int i = 0; i < (n); i++)
typedef long long ll;

// プロトタイプ宣言
extern void LoadTagCsv(string filename);

// 構造体宣言
struct Image {
    ll id;
    time_t date;
    double loc1;
    double loc2;
    int URL1;
    int URL2;
    string URL3;
};

struct Tag {
    string name;
    ll img[100];
};

// グローバル変数
extern vector<Image> images;
extern vector<Tag> tags;

#endif