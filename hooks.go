package main

// Register gamemode handlers in init(). The plugin ships with no gameplay logic.
//
// Example:
//
//	func init() {
//	    events.OnPlayerConnect = func(playerID int) {
//	        API.Player.SendMessage(playerID, ColourCyan, "Welcome!")
//	    }
//	    events.OnPlayerCommand = func(playerID int, cmd string) FilterResult {
//	        if cmd == "/pos" {
//	            pos := API.Player.Position(playerID)
//	            API.Player.SendMessage(playerID, ColourWhite, API.Player.FormatPosition(pos))
//	            return FilterDeny
//	        }
//	        return FilterAllow
//	    }
//	}

func init() {}
