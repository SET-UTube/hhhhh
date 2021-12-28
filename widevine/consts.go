package widevine

const DefaultPrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEArk8zHaSrt4Cx/XC6CDU1zDEj3izBIrEtBwa8kJsuh5fO3tol
JhGaarshuWT9qJMOpD6MPs6eJcywENVz4VN1v4KlS8rpWRbzIZ8Tl/oXj8rb2xnV
njqvjgB78Oc/OnPnbSUNlfhepMeIzIviuLMaWzhuNIcSZFrN6fjXqSzNpqO1roGI
QM1HB8UYgAd4hHF0ZEZYuLl5ja4HKcCV7eYMVAWyiGEorOGhFKKRfkXEmPruW9Rd
XWWZ9wEjf6Qlsaf66rdRBQOXEMWP2ChGSz3PQQ7wpA6YpzNCVIAd8OxoPBPakN+V
/13yR23Jyv8S7WzR7vV/TniufteDRH8sJvZMFQIDAQABAoIBABYcX0pLFWovVcSl
nD+GymZ2tEtOR4iIS4Mo6Fn6iVYzXE86AjzYPkt8Jdy/0gpkbxbIBV/BM1/tnUbh
YLPsM5NBtgiNenit19UDugdM9tirXaSnHAkYfiTn7FDkcImQYsKeUOEdcpn54qE5
vF57/6OjHp2kpaFbwGOFyIuB7qtMwqlcYX0vnOttJ0BWwKIIh5gi4yNbciqJKZ/5
P6WbKZVvD7rGsRGYroF35pPTy48cWVGYR1OokXZ35FRgX5QVM1bzP2Q0ExxGB5jT
7hCmA3DJLooKD2Aj9qX9v5ELMT+bhNIX3Q4aKX7xPZOcBPbugRf8Tuq8l4ThOinQ
6eLJaTECgYEA5gsW2SJh8eMV5mLjaW9K6HALhXkeVN4GDzDTLiaTx0VcHXnJC07X
nJLx0piw/mR1VoUlXHjfVbtnAN7mDJstqh/rBxhqQrCkXbS10qoRbIt/T+SuNf1s
WZcfMj0aPMwz9q2jkbfARV3z+fcJETtEbKDUdwm6ketIpOVwMzDVBvECgYEAwfo1
e+X2bzuEGfqmtkFvcrSSLg+OiEhsfgvS0Y2yJEYcUvw+RwH066DBx/tVX1kCFLX8
H0/EwIBycKb07bFBnj4gmbqbB2zpo65d3IErByQfuCzW1IymjwG+s+Tx9qGPfRvG
TxhoznOOuHO72Xb82B8pxe2/qqOAfIV4zlYJf2UCgYEA2ih0D5E6v5DyqNzo+4ku
ycXQN1EIgcVYi7lq3F57UMQnOlDPZyjq8rKsIGLrnyUX3ehA6TQ74GrroPjBw/y5
zpecZMszonEwPylsMQ9VnNGh99tPlyXxRfk5/YPSyQuC0BIVh9Bxx5b1E/3BnJTP
LBFNzGHujAlMiAyKXhGWRJECgYBFtmd8VKQhS7FpKMS7YX7tKWoTtbGS1vxuvE8S
0qrAEJZjWJYFLPXZrNeXyILhFnsB+HlYw3FBgagfRlFmDzs25LsQpJjMrV62XZcM
BTvygBAKP8shbj75zDW+LzyqV1vbKZ02ld4svCkBr05GlFXAUkrQAGbOq54kok9N
UGxvZQKBgQCGr0rGnCBmKYDUkqZm9FXPDcYWtbiud9zlsaW+xk3UR0fFHuGlK6rW
nxV9k5OOJsZL9MUZ4sy8Ob62ToxB23T+02E0QOgzN4tQjqHOUkMO5C/ErT8jFuhY
J9H1SW5VOZxxMW5VP+iOXmieAJ94CbjRSI26MGsE+6aeK3SsmSCGYg==
-----END RSA PRIVATE KEY-----
`

var DefaultClientID = []byte{8, 1, 18, 154, 11, 10, 221, 3, 8, 2, 18, 16, 207, 65, 222, 65, 106, 16, 14, 177, 84, 131, 98, 11, 160, 121, 101, 21, 24, 168, 223, 228, 133, 6, 34, 142, 2, 48, 130, 1, 10, 2, 130, 1, 1, 0, 174, 79, 51, 29, 164, 171, 183, 128, 177, 253, 112, 186, 8, 53, 53, 204, 49, 35, 222, 44, 193, 34, 177, 45, 7, 6, 188, 144, 155, 46, 135, 151, 206, 222, 218, 37, 38, 17, 154, 106, 187, 33, 185, 100, 253, 168, 147, 14, 164, 62, 140, 62, 206, 158, 37, 204, 176, 16, 213, 115, 225, 83, 117, 191, 130, 165, 75, 202, 233, 89, 22, 243, 33, 159, 19, 151, 250, 23, 143, 202, 219, 219, 25, 213, 158, 58, 175, 142, 0, 123, 240, 231, 63, 58, 115, 231, 109, 37, 13, 149, 248, 94, 164, 199, 136, 204, 139, 226, 184, 179, 26, 91, 56, 110, 52, 135, 18, 100, 90, 205, 233, 248, 215, 169, 44, 205, 166, 163, 181, 174, 129, 136, 64, 205, 71, 7, 197, 24, 128, 7, 120, 132, 113, 116, 100, 70, 88, 184, 185, 121, 141, 174, 7, 41, 192, 149, 237, 230, 12, 84, 5, 178, 136, 97, 40, 172, 225, 161, 20, 162, 145, 126, 69, 196, 152, 250, 238, 91, 212, 93, 93, 101, 153, 247, 1, 35, 127, 164, 37, 177, 167, 250, 234, 183, 81, 5, 3, 151, 16, 197, 143, 216, 40, 70, 75, 61, 207, 65, 14, 240, 164, 14, 152, 167, 51, 66, 84, 128, 29, 240, 236, 104, 60, 19, 218, 144, 223, 149, 255, 93, 242, 71, 109, 201, 202, 255, 18, 237, 108, 209, 238, 245, 127, 78, 120, 174, 126, 215, 131, 68, 127, 44, 38, 246, 76, 21, 2, 3, 1, 0, 1, 40, 195, 101, 72, 1, 82, 170, 1, 8, 1, 16, 0, 26, 129, 1, 4, 126, 185, 114, 76, 62, 145, 153, 81, 194, 228, 163, 151, 90, 56, 154, 9, 143, 171, 190, 177, 74, 229, 253, 3, 3, 148, 203, 92, 248, 175, 140, 190, 242, 143, 62, 238, 182, 178, 183, 39, 55, 20, 46, 95, 229, 186, 139, 216, 176, 22, 153, 57, 61, 66, 80, 151, 169, 79, 236, 197, 33, 136, 231, 246, 13, 226, 130, 44, 176, 5, 83, 149, 29, 14, 132, 89, 3, 202, 49, 103, 72, 120, 78, 28, 183, 57, 182, 110, 103, 121, 195, 214, 154, 192, 172, 100, 90, 63, 53, 164, 43, 187, 1, 118, 132, 152, 17, 254, 230, 146, 9, 70, 106, 169, 22, 190, 94, 60, 86, 147, 38, 232, 246, 148, 26, 162, 243, 181, 34, 32, 17, 32, 200, 120, 233, 63, 86, 167, 123, 157, 195, 31, 112, 147, 192, 104, 35, 178, 218, 188, 31, 37, 197, 73, 232, 208, 99, 142, 225, 186, 103, 177, 18, 128, 2, 163, 253, 199, 226, 90, 143, 6, 60, 70, 176, 5, 51, 145, 67, 220, 203, 183, 3, 151, 129, 185, 44, 189, 24, 143, 223, 198, 218, 12, 175, 244, 99, 234, 67, 37, 242, 235, 217, 175, 214, 84, 185, 9, 104, 36, 29, 74, 222, 188, 191, 223, 160, 126, 230, 22, 87, 89, 178, 0, 64, 128, 164, 221, 18, 94, 229, 179, 206, 244, 152, 223, 186, 203, 39, 47, 66, 2, 152, 144, 150, 110, 239, 174, 152, 216, 81, 188, 52, 241, 193, 82, 55, 29, 88, 56, 48, 57, 97, 101, 160, 252, 222, 119, 123, 238, 178, 11, 47, 77, 62, 13, 106, 166, 168, 34, 31, 128, 196, 132, 166, 121, 6, 47, 138, 157, 29, 82, 116, 37, 116, 103, 37, 54, 30, 51, 215, 141, 197, 225, 5, 127, 199, 44, 147, 142, 10, 241, 167, 73, 110, 51, 152, 132, 133, 124, 85, 16, 143, 26, 85, 203, 168, 28, 89, 194, 204, 189, 65, 222, 149, 17, 73, 79, 109, 101, 119, 248, 35, 126, 84, 58, 219, 88, 241, 180, 96, 208, 220, 25, 160, 139, 70, 237, 8, 98, 29, 162, 36, 245, 159, 118, 82, 170, 184, 8, 106, 237, 203, 21, 164, 177, 28, 239, 77, 208, 195, 86, 184, 20, 59, 19, 157, 86, 230, 5, 61, 205, 130, 34, 76, 45, 171, 189, 162, 230, 237, 90, 188, 142, 133, 21, 43, 246, 6, 147, 64, 244, 124, 204, 83, 225, 145, 74, 215, 73, 60, 26, 180, 5, 10, 174, 2, 8, 1, 18, 16, 32, 85, 111, 153, 7, 247, 235, 104, 53, 84, 52, 209, 89, 180, 9, 242, 24, 201, 183, 240, 225, 5, 34, 142, 2, 48, 130, 1, 10, 2, 130, 1, 1, 0, 203, 157, 139, 87, 226, 128, 226, 233, 229, 151, 125, 7, 201, 138, 22, 89, 202, 62, 157, 41, 188, 93, 13, 139, 6, 43, 119, 25, 87, 83, 122, 140, 100, 172, 106, 212, 162, 83, 190, 172, 77, 33, 156, 235, 101, 170, 174, 56, 15, 5, 199, 30, 124, 56, 209, 183, 46, 210, 53, 152, 163, 249, 123, 113, 55, 83, 102, 46, 29, 66, 34, 21, 99, 50, 29, 68, 237, 134, 244, 27, 153, 61, 80, 20, 74, 105, 52, 87, 244, 143, 173, 185, 35, 141, 143, 90, 26, 203, 153, 194, 24, 70, 193, 49, 158, 136, 169, 99, 52, 249, 171, 52, 220, 176, 188, 149, 100, 152, 204, 114, 95, 166, 85, 204, 35, 208, 13, 176, 23, 151, 94, 29, 101, 198, 66, 119, 41, 150, 39, 31, 82, 91, 113, 56, 119, 126, 129, 246, 118, 250, 197, 47, 58, 83, 228, 195, 192, 74, 55, 27, 196, 203, 20, 161, 235, 200, 139, 23, 131, 172, 92, 78, 223, 220, 56, 245, 97, 97, 233, 88, 118, 163, 63, 48, 168, 104, 155, 180, 33, 198, 19, 119, 241, 89, 235, 3, 209, 71, 81, 72, 57, 233, 3, 223, 200, 54, 5, 97, 185, 254, 121, 187, 92, 78, 250, 91, 127, 156, 67, 128, 153, 95, 50, 49, 18, 51, 124, 44, 230, 197, 79, 47, 49, 142, 69, 32, 250, 193, 216, 16, 63, 5, 162, 78, 99, 205, 23, 172, 88, 251, 118, 213, 51, 125, 35, 239, 2, 3, 1, 0, 1, 40, 195, 101, 18, 128, 3, 75, 119, 79, 75, 181, 36, 190, 54, 253, 232, 139, 246, 241, 19, 70, 151, 96, 124, 33, 126, 109, 35, 42, 102, 99, 245, 130, 98, 82, 50, 107, 167, 125, 193, 218, 110, 60, 213, 29, 218, 25, 106, 34, 113, 99, 71, 175, 15, 199, 146, 244, 129, 14, 255, 255, 167, 57, 68, 154, 193, 89, 49, 174, 182, 141, 14, 212, 58, 70, 15, 72, 61, 166, 202, 36, 170, 6, 64, 191, 94, 47, 128, 72, 170, 130, 51, 69, 123, 144, 139, 216, 58, 171, 190, 170, 30, 101, 232, 21, 234, 246, 61, 201, 109, 59, 255, 120, 45, 55, 89, 186, 8, 109, 235, 149, 172, 71, 4, 171, 184, 244, 52, 163, 243, 201, 5, 5, 111, 253, 193, 166, 226, 115, 11, 149, 96, 178, 107, 193, 179, 15, 109, 11, 93, 129, 5, 172, 170, 121, 46, 46, 82, 197, 140, 4, 162, 99, 232, 37, 239, 181, 96, 161, 31, 233, 43, 78, 196, 72, 239, 180, 218, 140, 138, 252, 251, 124, 216, 154, 205, 249, 150, 215, 191, 33, 132, 74, 155, 103, 91, 28, 51, 111, 98, 79, 90, 99, 11, 58, 191, 235, 72, 128, 46, 156, 128, 151, 81, 116, 198, 148, 11, 237, 212, 98, 129, 128, 202, 217, 100, 122, 235, 246, 241, 176, 88, 51, 56, 160, 29, 7, 55, 195, 138, 212, 19, 183, 71, 65, 161, 254, 135, 253, 153, 195, 49, 18, 2, 63, 56, 140, 31, 120, 174, 149, 154, 167, 83, 137, 120, 137, 237, 205, 49, 96, 70, 82, 59, 2, 47, 147, 241, 137, 226, 238, 254, 220, 173, 79, 37, 185, 246, 47, 1, 219, 203, 11, 3, 244, 90, 39, 117, 201, 134, 52, 14, 173, 56, 252, 21, 154, 96, 41, 217, 54, 207, 144, 22, 29, 194, 197, 214, 211, 245, 208, 1, 105, 183, 16, 234, 87, 24, 201, 158, 251, 122, 138, 206, 62, 123, 78, 98, 23, 119, 142, 35, 86, 22, 216, 134, 224, 130, 125, 8, 119, 57, 121, 87, 33, 98, 158, 249, 151, 195, 162, 109, 204, 73, 77, 144, 191, 85, 54, 201, 76, 138, 179, 235, 192, 249, 131, 250, 192, 241, 241, 152, 205, 237, 59, 227, 8, 109, 87, 54, 26, 22, 10, 12, 99, 111, 109, 112, 97, 110, 121, 95, 110, 97, 109, 101, 18, 6, 76, 69, 78, 79, 86, 79, 26, 29, 10, 10, 109, 111, 100, 101, 108, 95, 110, 97, 109, 101, 18, 15, 76, 101, 110, 111, 118, 111, 32, 84, 66, 45, 88, 53, 48, 53, 88, 26, 30, 10, 17, 97, 114, 99, 104, 105, 116, 101, 99, 116, 117, 114, 101, 95, 110, 97, 109, 101, 18, 9, 97, 114, 109, 54, 52, 45, 118, 56, 97, 26, 29, 10, 11, 100, 101, 118, 105, 99, 101, 95, 110, 97, 109, 101, 18, 14, 76, 101, 110, 111, 118, 111, 84, 66, 45, 88, 53, 48, 53, 88, 26, 30, 10, 12, 112, 114, 111, 100, 117, 99, 116, 95, 110, 97, 109, 101, 18, 14, 76, 101, 110, 111, 118, 111, 84, 66, 45, 88, 53, 48, 53, 88, 26, 91, 10, 10, 98, 117, 105, 108, 100, 95, 105, 110, 102, 111, 18, 77, 76, 101, 110, 111, 118, 111, 47, 76, 101, 110, 111, 118, 111, 84, 66, 45, 88, 53, 48, 53, 88, 47, 84, 66, 45, 88, 53, 48, 53, 88, 58, 57, 47, 80, 75, 81, 49, 46, 49, 55, 49, 49, 48, 50, 46, 48, 48, 50, 47, 48, 52, 56, 95, 49, 56, 48, 50, 49, 49, 58, 117, 115, 101, 114, 47, 114, 101, 108, 101, 97, 115, 101, 45, 107, 101, 121, 115, 26, 45, 10, 9, 100, 101, 118, 105, 99, 101, 95, 105, 100, 18, 32, 84, 66, 45, 88, 53, 48, 53, 88, 48, 48, 48, 50, 48, 50, 55, 53, 53, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 26, 30, 10, 20, 119, 105, 100, 101, 118, 105, 110, 101, 95, 99, 100, 109, 95, 118, 101, 114, 115, 105, 111, 110, 18, 6, 49, 51, 46, 48, 46, 48, 26, 36, 10, 31, 111, 101, 109, 95, 99, 114, 121, 112, 116, 111, 95, 115, 101, 99, 117, 114, 105, 116, 121, 95, 112, 97, 116, 99, 104, 95, 108, 101, 118, 101, 108, 18, 1, 48, 50, 12, 16, 1, 32, 4, 40, 12, 48, 1, 64, 0, 72, 1}
