package cmd

const (
	rootComponentDirTemplate = `[){[toImports .RestImports]}(]
export { [){[toExports .RestExports]}(][){[.Name]}(] };`

	newComponentTemplate = `import React, { Component } from 'react';
import styles from './[){[.Name]}(].css';

class [){[.Name]}(] extends Component {
	render() {
		return (
			
		)
	}
}

export default [){[.Name]}(];`

	newContainerTemplate = `import React, { Component } from 'react';
import styles from './[){[.Name]}(].css';

class [){[.Name]}(] extends Component {
	render() {
		return(
			<div>{'[){[.Name]}(] works'}</div>
		)
	}
}

export default [){[.Name]}(];`
)
