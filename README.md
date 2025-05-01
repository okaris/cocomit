# cocomit

**Estimate the cost of every commit in your Git history using COCOMO.**

<img width="622" alt="image" src="https://github.com/user-attachments/assets/1f98f24a-77ca-4f10-b2d8-d60c34ea3b75" />

---

## Features

- Parses Git history and diffs
- Calculates per-commit cost based on COCOMO (organic mode)
- No dependencies beyond Git and Go

---

## COCOMO Assumptions

- **Model**: Basic COCOMO (Organic)
- **Effort**: `Effort = 2.4 × (SLOC / 1000)^1.05`
- **EAF**: 1.0 (nominal)
- **Annual Developer Salary**: $120,000
- **Overhead Multiplier**: 1.3 (benefits, infra, admin)
- **Cost**: `Cost = Effort × (Salary / 12) × Overhead`

---

## Installation

```bash
git clone https://github.com/YOUR-USER/cocomit.git
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

⸻

### Output Example

```bash
Assumptions:
Model:        Basic COCOMO (Organic)
Effort:       Effort = 2.4 × (SLOC / 1000)^1.05
EAF:          1.0 (nominal)
Annual Wage:  $120,000
Overhead:     1.3

d0fc851cd302 — refactor parser logic — $98,677.69
265f3027f28f — fix spacing — $1,854.77
```

⸻

License

MIT
