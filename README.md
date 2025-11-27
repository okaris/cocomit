# cocomit

**Estimate the cost of every commit in your Git history.**

<img width="622" alt="image" src="https://github.com/user-attachments/assets/1f98f24a-77ca-4f10-b2d8-d60c34ea3b75" />

---

## Features

- Parses Git history and diffs
- Calculates per-commit cost based on lines of code
- Shows total LoC and cost for the entire repo
- No dependencies beyond Git and Go

---

## Assumptions

- **Model**: Single Developer Linear Estimate
- **Effort**: `Effort = (SLOC / 50) hours`
- **EAF**: 1.0 (nominal)
- **Hourly Wage**: $65.79 ($120K/year)
- **Overhead Multiplier**: 1.3 (benefits, infra, admin)
- **Cost**: `Cost = Effort × Hourly Wage × Overhead`

---

## Installation

```bash
git clone https://github.com/okaris/cocomit.git
cd cocomit
go build -o cocomit
sudo mv cocomit /usr/local/bin/
```

---

## Usage

Run inside any Git repo:

```bash
cocomit
```

Show only totals (no per-commit output):

```bash
cocomit -total
```

---

### Output Example

```bash
#cocomit
Assumptions:
- Model:        Single Developer Linear Estimate
- Effort:       Effort = (SLOC / 50) hours → PM = hours / 152
- EAF:          1.0 (nominal)
- Hourly Wage:  $65.79 ($120K/year)
- Overhead:     1.3 (benefits, infra, etc.)
- Cost:         Effort × Hours × Hourly Wage × Overhead

d0fc851 | refactor parser logic | $98.67
265f302 | fix spacing | $18.54

────────────────────────────────────────
Total LoC:  12345
Total Cost: $1,234.56
```

With `-total` flag:

```bash
#cocomit
Total LoC:  12345
Total Cost: $1,234.56
```

---

## License

MIT
