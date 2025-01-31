public class Solution {
    public int CharacterReplacement(string s, int k) {
     int count = 0;
      Dictionary<char, int> map = new Dictionary<char, int>();
       int l = 0;
        int maxfrequancy = 0; 
         for (int r = 0; r < s.Length;r++)
         {
              if (map.ContainsKey(s[r])) map[s[r]]++;
              else map[s[r]] = 1;

               maxfrequancy = Math.Max(maxfrequancy, map[s[r]]);

              while (r - l + 1 - maxfrequancy > k)
                {
                   map[s[l]]--;
                   l++;
                }
             count = Math.Max(count, r - l + 1);
         }

          return count;    
    }
}