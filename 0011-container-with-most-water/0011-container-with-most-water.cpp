class Solution {
public:
    int maxArea(vector<int>& height) {
        int res = 0;
        int n = height.size();
        int left = 0, right = n-1;

        while(left < right && left < n-1 && right > 0){
            int area = (right - left) * min(height[left], height[right]);
            res = max(res,area);

            if(height[left] < height[right]) left++;
            else right--;
        }
        return res;
    }
};