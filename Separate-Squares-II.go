1func separateSquares(squares [][]int) float64 {
2    type Event struct {
3        y, typ, x1, x2 int
4    }
5    
6    events := make([]Event, 0, len(squares)*2)
7    for _, sq := range squares {
8        x, y, l := sq[0], sq[1], sq[2]
9        events = append(events, Event{y, 1, x, x + l})
10        events = append(events, Event{y + l, -1, x, x + l})
11    }
12    
13    sort.Slice(events, func(i, j int) bool {
14        if events[i].y != events[j].y {
15            return events[i].y < events[j].y
16        }
17        return events[i].typ > events[j].typ
18    })
19    
20    type Area struct {
21        y, h, w int
22    }
23    
24    areas := make([]Area, 0)
25    xs := make([][2]int, 0)
26    prevY := events[0].y
27    total := 0
28    
29    for _, ev := range events {
30        if ev.y > prevY && len(xs) > 0 {
31            h := ev.y - prevY
32            w := unionLen(xs)
33            areas = append(areas, Area{prevY, h, w})
34            total += h * w
35        }
36        
37        if ev.typ == 1 {
38            xs = append(xs, [2]int{ev.x1, ev.x2})
39        } else {
40            for i := 0; i < len(xs); i++ {
41                if xs[i][0] == ev.x1 && xs[i][1] == ev.x2 {
42                    xs = append(xs[:i], xs[i+1:]...)
43                    break
44                }
45            }
46        }
47        
48        prevY = ev.y
49    }
50    
51    half := float64(total) / 2.0
52    acc := 0.0
53    
54    for _, a := range areas {
55        area := float64(a.h * a.w)
56        if acc+area >= half {
57            return float64(a.y) + (half-acc)/float64(a.w)
58        }
59        acc += area
60    }
61    
62    return 0.0
63}
64
65func unionLen(intervals [][2]int) int {
66    if len(intervals) == 0 {
67        return 0
68    }
69    
70    sorted := make([][2]int, len(intervals))
71    copy(sorted, intervals)
72    
73    sort.Slice(sorted, func(i, j int) bool {
74        return sorted[i][0] < sorted[j][0]
75    })
76    
77    res := 0
78    end := -1000000000
79    
80    for _, interval := range sorted {
81        a, b := interval[0], interval[1]
82        if a > end {
83            res += b - a
84            end = b
85        } else if b > end {
86            res += b - end
87            end = b
88        }
89    }
90    
91    return res
92}