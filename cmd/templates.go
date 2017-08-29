package cmd

const (
	rootComponentDirTemplate = `[){[toImports .RestImports]}(]
export { [){[toExports .RestExports]}(][){[.Name]}(] };`

	newComponentTemplate = `import React, { Component } from 'react';
import styles from './[){[.Name]}(].sass';

class [){[.Name]}(] extends Component {
	render() {
		return (
			<div>[){[.Name]}(] component</div>
		)
	}
}

export default [){[.Name]}(];`
)
