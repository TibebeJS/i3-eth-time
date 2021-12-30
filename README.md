# i3-eth-time

Display ethiopian calendar date/time in i3-wm!
Intended for use in i3-status.

### How to use
- Clone/Download repo and compile
- Open up your `i3config` file (usually located at `/home/your-name/.config/i3/config`)
- Look for the line that starts with `status_command` inside `bar` configuration
- Replace with "`status_command exec bin-path`" (replace "`bin-path`" with a location to the binary you have compiled in step 1)
 
### Example

```i3
bar {
  status_command exec /home/tibebe/.config/i3-eth-time # NOTE: replace with where your binary is located at.
}
```
