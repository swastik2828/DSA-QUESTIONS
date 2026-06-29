class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        vector<int> ls(256, -1);
        int l = 0, r = 0, maxlen = 0;
        int n =  s.size();

        while(r < n){
            if(ls[s[r]] != -1){
                if(ls[s[r]] >= l){
                    l = ls[s[r]] + 1;
                }
            }
            int len = r-l+1;
            maxlen = max(len,maxlen);

            ls[s[r]] = r;
            r++;
        }
        return maxlen;
    }
};