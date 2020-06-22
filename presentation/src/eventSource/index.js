import store from '@/store'

const Handlers = [
    activeGames,
    openGames,
    currentGame,
    notifications
];

export default function configureEventSources () {
    const eventSource = new EventSource(ApiRoutes.EventSource)
    for (const handler of Handlers) {
      eventSource.addEventListener(handler.eventType, event => {
        handler.handle(event)
      })
    }
  }

const activeGames = {
    eventType:'Update',
    handle: function(event) {
        store.commit('updateActiveGamesList', [])
    }
};

const openGames = {
    eventType:'NODE_STATUS',
    handle: function(event) {
        store.commit('updateOpenGamesList', [])
    }
};

const currentGame = {
    eventType:'NODE_STATUS',
    handle: function(event) {
        store.commit('updateActiveGamesList', [])
    }
};

const notifications = {
    eventType:'NODE_STATUS',
    handle: function(event) {
        store.commit('updateActiveGamesList', [])
    }
};