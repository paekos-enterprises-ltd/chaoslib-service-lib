# Chaoslib Service API Documentation

## Random Number Generation

**Endpoints:** `/api/rand/...`

- `/api/rand/uint64`
- `/api/rand/int63`
- `/api/rand/uint32`
- `/api/rand/int31`
- `/api/rand/int`
- `/api/rand/float64`
- `/api/rand/float32`
- `/api/rand/int63n?n=100` (takes 'n' parameter)
- `/api/rand/int31n?n=100` (takes 'n' parameter)
- `/api/rand/intn?n=100` (takes 'n' parameter)
- `/api/rand/perm?n=5` (takes 'n' parameter)

**Example Response (for /api/rand/uint64):**
```json
{
    "result_uint64": 1234567890123456789
}
```

## Crash Multiplier

**Endpoint:** `/api/crash`

**Example Response:**
```json
{
    "result_int64": 150
}
```

## Weight

**Endpoint:** `/api/weight`

Takes an optional `count` parameter.

**Example Response (single):**
```json
{
    "result_int64": 1234567
}
```

**Example Response (count=5):**
```json
{
    "result_int64_array": [1234567, 8765432, 3456789, 9876543, 4567890]
}
```

## Slots

**Endpoint:** `/api/slots`

Requires a `reels` parameter (comma-separated integers). Takes an optional `count` parameter.

**Example Response (single, reels=10,20,30):**
```json
{
    "result_int_array": [5, 15, 25]
}
```

**Example Response (count=3, reels=10,20,30):**
```json
{
    "result_int2_array": [[5, 15, 25], [1, 2, 3], [9, 18, 29]]
}
```

## Penalty

**Endpoint:** `/api/penalty`

**Example Response:**
```json
{
    "result_int32": 12
}
```
