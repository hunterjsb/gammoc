package client

type Endpoint string

const (
	// My characters
	EndpointActionMove             Endpoint = "/my/{name}/action/move"
	EndpointActionEquipItem        Endpoint = "/my/{name}/action/equip/"
	EndpointActionUnequipItem      Endpoint = "/my/{name}/action/unequip/"
	EndpointActionFight            Endpoint = "/my/{name}/action/fight"
	EndpointActionGathering        Endpoint = "/my/{name}/action/gathering"
	EndpointActionCrafting         Endpoint = "/my/{name}/action/crafting"
	EndpointActionDepositBank      Endpoint = "/my/{name}/action/bank/deposit"
	EndpointActionDepositBankGold  Endpoint = "/my/{name}/action/bank/deposit/gold"
	EndpointActionRecycling        Endpoint = "/my/{name}/action/recycling"
	EndpointActionWithdrawBank     Endpoint = "/my/{name}/action/bank/withdraw"
	EndpointActionWithdrawBankGold Endpoint = "/my/{name}/action/bank/withdraw/gold"
	EndpointActionGEBuy            Endpoint = "/my/{name}/action/ge/buy"
	EndpointActionGESell           Endpoint = "/my/{name}/action/ge/sell"
	EndpointActionAcceptNewTask    Endpoint = "/my/{name}/action/task/new"
	EndpointActionCompleteTask     Endpoint = "/my/{name}/action/task/complete"
	EndpointActionTaskExchange     Endpoint = "/my/{name}/action/task/exchange"
	EndpointActionDeleteItem       Endpoint = "/my/{name}/action/delete"
	EndpointGetCharacterLogs       Endpoint = "/my/{name}/logs"
	EndpointGetAllCharactersLogs   Endpoint = "/my/logs"
	EndpointGetMyCharacters        Endpoint = "/my/characters"

	// My account
	EndpointGetBankItems   Endpoint = "/my/bank/items"
	EndpointGetBankGold    Endpoint = "/my/bank/gold"
	EndpointChangePassword Endpoint = "/my/change_password"

	// Characters
	EndpointCreateCharacter  Endpoint = "/characters/create"
	EndpointGetAllCharacters Endpoint = "/characters/"
	EndpointGetCharacter     Endpoint = "/characters/{name}"

	// Maps
	EndpointGetAllMaps Endpoint = "/maps/"
	EndpointGetMap     Endpoint = "/maps/{x}/{y}"

	// Items
	EndpointGetAllItems Endpoint = "/items/"
	EndpointGetItem     Endpoint = "/items/{code}"

	// Monsters
	EndpointGetAllMonsters Endpoint = "/monsters/"
	EndpointGetMonster     Endpoint = "/monsters/{code}"

	// Resources
	EndpointGetAllResources Endpoint = "/resources/"
	EndpointGetResource     Endpoint = "/resources/{code}"

	// Events
	EndpointGetAllEvents Endpoint = "/events/"

	// Grand Exchange
	EndpointGetAllGEItems Endpoint = "/ge/"
	EndpointGetGEItem     Endpoint = "/ge/{code}"

	// Accounts
	EndpointCreateAccount Endpoint = "/accounts/create"

	// Token
	EndpointGenerateToken Endpoint = "/token/"

	// Status
	EndpointGetStatus Endpoint = "/"
)
