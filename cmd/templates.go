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

	newReducerTemplate = `const initialState = {
	[){[.Name]}(]: undefined
}

export default (state=initialState, action) => {
	switch(action.type) {
		default:
			return state
	}
	return state
}`

	newActionTemplate = `export const [){[.Name]}(] = () => ({})`
)
