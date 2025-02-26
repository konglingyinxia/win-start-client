const state = {
    isCollapseTag: false,
    selectMenu: []
}
const mutations = {
    collapseTag(state: any) {
        state.isCollapseTag = !state.isCollapseTag
    },
    addMenu(state: any, payload: any) {
        if (state.selectMenu.findIndex(item => item.path === payload.path) === -1) {
            state.selectMenu.push(payload)
        }
    },
    removeTab(state: any, payload: any) {
        const index = state.selectMenu.findIndex(item => item.path === payload.path)
        state.selectMenu.splice(index, 1)
    }

}
export default {
    state,
    mutations
}
