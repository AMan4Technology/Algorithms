package GenerateParentheses

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	return validPathsByDFS(n, n, "", []string{})
}

func validPathsByDFS(left, right int, path string, paths []string) []string {
	if left == 0 && right == 0 {
		paths = append(paths, path)
		return paths
	}
	if left > 0 {
		paths = validPathsByDFS(left-1, right, path+"(", paths)
	}
	if left < right {
		paths = validPathsByDFS(left, right-1, path+")", paths)
	}
	return paths
}
