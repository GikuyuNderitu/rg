package cmd

const (
	rootComponentDirTemplate = `[){[toImports .RestImports]}(]
export { [){[toExports .RestExports]}(][){[.Name]}(] };`

	newComponentTemplate = `import React, { Component } from 'react';
import styles from './[){[.Name]}(].sass';

class [){[.Name]}(] extends Component {
	render() {
		
	}
}

export default [){[.Name]}(];`

	newContainerTemplate = `import React, { Component } from 'react';
import styles from './[){[.Name]}(].sass';

class [){[.Name]}(] extends Component {
	render() {
		
	}
}

export default [){[.Name]}(];`
)
