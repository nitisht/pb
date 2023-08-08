// Copyright (c) 2023 Cloudnatively Services Pvt Ltd
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package selection

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	items        []string
	focusIndex   int
	focus        bool
	FocusStyle   lipgloss.Style
	BlurredStyle lipgloss.Style
}

func (m *Model) Focus() tea.Cmd {
	m.focus = true
	return nil
}

func (m *Model) Blur() {
	m.focus = false
}

func (m *Model) Focused() bool {
	return m.focus
}

func (m *Model) Value() string {
	return m.items[m.focusIndex]
}

func New(items []string) Model {
	m := Model{
		focusIndex: 0,
		focus:      false,
		items:      items,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyLeft:
			if m.focusIndex > 0 {
				m.focusIndex -= 1
			}
		case tea.KeyRight:
			if m.focusIndex < len(m.items)-1 {
				m.focusIndex += 1
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	render := make([]string, len(m.items))

	for idx, item := range m.items {
		if idx == m.focusIndex {
			render[idx] = m.FocusStyle.Render(item)
		} else {
			render[idx] = m.BlurredStyle.Render(item)
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, render...)
}