1class Solution(object):
2    def separateSquares(self, squares):
3        total_area = 0
4        events = []
5        
6        for x, y, l in squares:
7            total_area += l * l
8            events.append((y, l))      
9            events.append((y + l, -l))  
10        
11        events.sort()
12        target_area = total_area / 2.0
13        
14        current_area = 0
15        current_width = 0
16        prev_y = events[0][0]
17        
18        for curr_y, width_change in events:
19            height_diff = curr_y - prev_y
20            area_increase = current_width * height_diff
21            
22            if current_area + area_increase >= target_area:
23                needed_area = target_area - current_area
24                return prev_y + needed_area / current_width
25            
26            current_area += area_increase
27            current_width += width_change
28            prev_y = curr_y
29        
30        return 0.0