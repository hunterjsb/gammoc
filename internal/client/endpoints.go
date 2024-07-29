package client

type Endpoint string

const (
	// My characters
	// `%s` is name
	EndpointActionMove             Endpoint = "/my/%s/action/move"
	EndpointActionEquipItem        Endpoint = "/my/%s/action/equip/"
	EndpointActionUnequipItem      Endpoint = "/my/%s/action/unequip/"
	EndpointActionFight            Endpoint = "/my/%s/action/fight"
	EndpointActionGathering        Endpoint = "/my/%s/action/gathering"
	EndpointActionCrafting         Endpoint = "/my/%s/action/crafting"
	EndpointActionDepositBank      Endpoint = "/my/%s/action/bank/deposit"
	EndpointActionDepositBankGold  Endpoint = "/my/%s/action/bank/deposit/gold"
	EndpointActionRecycling        Endpoint = "/my/%s/action/recycling"
	EndpointActionWithdrawBank     Endpoint = "/my/%s/action/bank/withdraw"
	EndpointActionWithdrawBankGold Endpoint = "/my/%s/action/bank/withdraw/gold"
	EndpointActionGEBuy            Endpoint = "/my/%s/action/ge/buy"
	EndpointActionGESell           Endpoint = "/my/%s/action/ge/sell"
	EndpointActionAcceptNewTask    Endpoint = "/my/%s/action/task/new"
	EndpointActionCompleteTask     Endpoint = "/my/%s/action/task/complete"
	EndpointActionTaskExchange     Endpoint = "/my/%s/action/task/exchange"
	EndpointActionDeleteItem       Endpoint = "/my/%s/action/delete"
	EndpointGetCharacterLogs       Endpoint = "/my/%s/logs"
	EndpointGetAllCharactersLogs   Endpoint = "/my/logs"
	EndpointGetMyCharacters        Endpoint = "/my/characters"

	// My account
	EndpointGetBankItems   Endpoint = "/my/bank/items"
	EndpointGetBankGold    Endpoint = "/my/bank/gold"
	EndpointChangePassword Endpoint = "/my/change_password"

	// Characters
	EndpointCreateCharacter  Endpoint = "/characters/create"
	EndpointGetAllCharacters Endpoint = "/characters/"
	EndpointGetCharacter     Endpoint = "/characters/%s"

	// Maps
	EndpointGetAllMaps Endpoint = "/maps/"
	EndpointGetMap     Endpoint = "/maps/%s/%s" // x, y

	// Items
	EndpointGetAllItems Endpoint = "/items/"
	EndpointGetItem     Endpoint = "/items/%s" // item code

	// Monsters
	EndpointGetAllMonsters Endpoint = "/monsters/"
	EndpointGetMonster     Endpoint = "/monsters/%s" // item code

	// Resources
	EndpointGetAllResources Endpoint = "/resources/"
	EndpointGetResource     Endpoint = "/resources/%s" // item code

	// Events
	EndpointGetAllEvents Endpoint = "/events/"

	// Grand Exchange
	EndpointGetAllGEItems Endpoint = "/ge/"
	EndpointGetGEItem     Endpoint = "/ge/%s" // item code

	// Accounts
	EndpointCreateAccount Endpoint = "/accounts/create"

	// Token
	EndpointGenerateToken Endpoint = "/token/"

	// Status
	EndpointGetStatus Endpoint = "/"
)
