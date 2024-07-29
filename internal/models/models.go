package models

import "time"

// Common response wrapper
type Response struct {
	Data interface{} `json:"data"`
}

// Character
type Character struct {
	Name               string    `json:"name"`
	Skin               string    `json:"skin"`
	Level              int       `json:"level"`
	XP                 int       `json:"xp"`
	MaxXP              int       `json:"max_xp"`
	TotalXP            int       `json:"total_xp"`
	Gold               int       `json:"gold"`
	X                  int       `json:"x"`
	Y                  int       `json:"y"`
	Cooldown           int       `json:"cooldown"`
	CooldownExpiration time.Time `json:"cooldown_expiration"`
	// Add other character fields as needed
}

// Map
type Map struct {
	Name    string     `json:"name"`
	Skin    string     `json:"skin"`
	X       int        `json:"x"`
	Y       int        `json:"y"`
	Content MapContent `json:"content"`
}

type MapContent struct {
	Monsters  []string `json:"monsters"`
	Resources []string `json:"resources"`
	// Add other content types as needed
}

// Movement
type MovementRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MovementResponse struct {
	Cooldown    CooldownInfo `json:"cooldown"`
	Destination Map          `json:"destination"`
	Character   Character    `json:"character"`
}

// Combat
type CombatResponse struct {
	Cooldown  CooldownInfo `json:"cooldown"`
	Fight     FightResult  `json:"fight"`
	Character Character    `json:"character"`
}

type FightResult struct {
	XP     int      `json:"xp"`
	Gold   int      `json:"gold"`
	Drops  []Drop   `json:"drops"`
	Turns  int      `json:"turns"`
	Logs   []string `json:"logs"`
	Result string   `json:"result"`
}

type Drop struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

// Gathering
type GatheringResponse struct {
	Cooldown  CooldownInfo `json:"cooldown"`
	Details   SkillInfo    `json:"details"`
	Character Character    `json:"character"`
}

type SkillInfo struct {
	XP    int    `json:"xp"`
	Items []Drop `json:"items"`
}

// Crafting
type CraftingRequest struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type CraftingResponse struct {
	Cooldown  CooldownInfo `json:"cooldown"`
	Details   SkillInfo    `json:"details"`
	Character Character    `json:"character"`
}

// Equipment
type EquipRequest struct {
	Code string `json:"code"`
	Slot string `json:"slot"`
}

type EquipResponse struct {
	Cooldown  CooldownInfo `json:"cooldown"`
	Slot      string       `json:"slot"`
	Item      Item         `json:"item"`
	Character Character    `json:"character"`
}

type Item struct {
	Name        string   `json:"name"`
	Code        string   `json:"code"`
	Level       int      `json:"level"`
	Type        string   `json:"type"`
	Subtype     string   `json:"subtype"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
}

type Effect struct {
	// Define effect structure
}

// Cooldown info
type CooldownInfo struct {
	TotalSeconds     int       `json:"totalSeconds"`
	RemainingSeconds int       `json:"remainingSeconds"`
	Expiration       time.Time `json:"expiration"`
	Reason           string    `json:"reason"`
}

// Bank
type BankItemsResponse struct {
	Data  []SimpleItem `json:"data"`
	Total int          `json:"total"`
	Page  int          `json:"page"`
	Size  int          `json:"size"`
	Pages int          `json:"pages"`
}

type SimpleItem struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

// Grand Exchange
type GEItemResponse struct {
	Item  string `json:"item"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
}

// Status
type StatusResponse struct {
	Status           string         `json:"status"`
	Version          string         `json:"version"`
	CharactersOnline int            `json:"characters_online"`
	Announcements    []Announcement `json:"announcements"`
	LastWipe         string         `json:"last_wipe"`
	NextWipe         string         `json:"next_wipe"`
}

type Announcement struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
