package org.treewalking.repeat.jira;

import org.junit.Assert;
import org.junit.Test;
import java.util.List;

public class JiraTest {

    @Test
    public void testCanExtractListOfProjectsFromJira() {
        // given
        Jira jira = givenAJiraEnvironmentWithASetOfProjects();
        // when
        List<Project> projects = whenIQueryForTheListOfProjects(jira);
        // then
        Assert.assertEquals(1, projects.size());
    }

    private Jira givenAJiraEnvironmentWithASetOfProjects() {
        return  new Jira("https://treewalking.atlassian.net/rest/api/3/project");
    }

    private List<Project> whenIQueryForTheListOfProjects(final Jira jira) {
        return jira.getProjects();
    }
}