package investsdk

import solana "github.com/gagliardetto/solana-go"

type OrcaSwapDirection int

const (
	AtoB OrcaSwapDirection = 0
	BtoA OrcaSwapDirection = 1
)

type OrcaSwapPool string

const (
	SBR_USDC   OrcaSwapPool = "CS7fA5n4c2D82dUoHrYzS3gAqgqaoVSfgsr18kitp2xo"
	SUNNY_USDC OrcaSwapPool = "GHuoeq9UnFBsBhMwH43eL3RWX5XVXbSRYJymmyMYpT7n"
)

type OperationRequest struct {
	Accounts map[string]interface{} `json:"accounts"`
	Data     map[string]interface{} `json:"data"`
	Platform string                 `json:"platform"`
}

type OperationResponse struct {
	Instructions        []Instruction  `json:"instructions,omitempty"`
	PrepareInstructions []Instruction  `json:"prepareInstructions,omitempty"`
	CleanupInstructions []Instruction  `json:"cleanupInstructions,omitempty"`
	SetupAccounts       []SetupAccount `json:"setupAccounts,omitempty"`
}

func (r *OperationResponse) FirstInstructionData() []byte {
	return r.Instructions[0].DataUnsafe
}

func (r *OperationResponse) SecondInstructionData() []byte {
	return r.Instructions[1].DataUnsafe
}

func (r *OperationResponse) AllPrepareInstructions() (instructions []solana.Instruction) {
	for _, i := range r.PrepareInstructions {
		instructions = append(instructions, i)
	}
	return
}

func (r *OperationResponse) MajorInstructions() (instructions []solana.Instruction) {
	for _, i := range r.Instructions {
		instructions = append(instructions, i)
	}
	return
}

func (r *OperationResponse) AllCleanupInstructions() (instructions []solana.Instruction) {
	for _, i := range r.CleanupInstructions {
		instructions = append(instructions, i)
	}
	return
}

func (r *OperationResponse) AllInstructions() (all []solana.Instruction) {
	for _, i := range r.PrepareInstructions {
		all = append(all, i)
	}
	for _, i := range r.Instructions {
		all = append(all, i)
	}
	for _, i := range r.CleanupInstructions {
		all = append(all, i)
	}
	return
}

var _ solana.Instruction = (*Instruction)(nil)

type Instruction struct {
	ProgramIDUnsafe solana.PublicKey `json:"programId,omitempty"`
	Keys            []AccountMeta    `json:"keys,omitempty"`
	DataUnsafe      []byte           `json:"data"`
}

type SetupAccount struct {
	Seed       string           `json:"seed,omitempty"`
	Pubkey     solana.PublicKey `json:"pubkey,omitempty"`
	Size       int32            `json:"size"`
	ProgramId  solana.PublicKey `json:"programId"`
	MinBalance int64            `json:"minBalance"`
}

type AccountMeta struct {
	PublicKey  solana.PublicKey `json:"pubkey,omitempty"`
	IsWritable bool             `json:"isWritable,omitempty"`
	IsSigner   bool             `json:"isSigner,omitempty"`
}

func (inst Instruction) ProgramID() solana.PublicKey {
	return inst.ProgramIDUnsafe
}

func (inst Instruction) Accounts() (accounts []*solana.AccountMeta) {
	for _, act := range inst.Keys {
		accounts = append(accounts, &solana.AccountMeta{
			PublicKey:  act.PublicKey,
			IsWritable: act.IsWritable,
			IsSigner:   act.IsSigner,
		})
	}
	return accounts
}

func (inst Instruction) Data() ([]byte, error) {
	return inst.DataUnsafe, nil
}
