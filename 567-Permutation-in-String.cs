using System.Diagnostics;
public class Solution {
     public bool CheckInclusion(string s1, string s2) {
        int[] count1 = new int[26];
        foreach(char c in s1 ){
                count1[c-'a']++;
        }
        int[] count2 = new int[26];
        int l = 0 ;
        int r = 0 ;
        while(r < s2.Length){
                count2[s2[r]-'a']++;
                if(r-l+1 > s1.Length){
                    count2[s2[l]-'a']--;
                    l++;
                }
                if(Enumerable.SequenceEqual(count1, count2)){
                    return true;
                }
                r++;
        }
        return false;
    }
}

      